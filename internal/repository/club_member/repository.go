package club_member

import (
	"context"
	"errors"
	"github.com/alitvinenko/ecareer_bot/internal/lib/e"
	"github.com/alitvinenko/ecareer_bot/internal/model"
	def "github.com/alitvinenko/ecareer_bot/internal/repository"
	repoModel "github.com/alitvinenko/ecareer_bot/internal/repository/club_member/model"
	"github.com/alitvinenko/ecareer_bot/internal/repository/converter"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"sync"
)

var _ def.ClubMemberRepository = (*repository)(nil)

type repository struct {
	db *gorm.DB
	m  sync.RWMutex
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, member *model.ClubMember) error {
	r.m.Lock()
	defer r.m.Unlock()

	if err := r.db.Debug().Create(converter.ToRepoFromClubMember(member)).Error; err != nil {
		return e.Wrap("error on create club member", err)
	}

	return nil
}

func (r *repository) Update(ctx context.Context, member *model.ClubMember) error {
	r.m.Lock()
	defer r.m.Unlock()

	err := r.db.Session(&gorm.Session{FullSaveAssociations: true}).Debug().Clauses(clause.OnConflict{UpdateAll: true}).Updates(converter.ToRepoFromClubMember(member)).Error
	if err != nil {
		return e.Wrap("error on update club member", err)
	}

	return nil
}

func (r *repository) Get(ctx context.Context, ID int) (*model.ClubMember, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	var clubMember repoModel.ClubMember
	err := r.db.Debug().Preload("Profile").First(&clubMember, "id = ?", ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return converter.ToClubMemberFromRepo(&clubMember), nil
}

func (r *repository) GetByUsername(ctx context.Context, username string) (*model.ClubMember, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	var clubMember repoModel.ClubMember
	err := r.db.Debug().Preload("Profile").First(&clubMember, "username = ?", username).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return converter.ToClubMemberFromRepo(&clubMember), nil
}
