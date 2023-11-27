package club_member

import (
	"context"
	"errors"
	"github.com/alitvinenko/ecareer_bot/internal/model"
	def "github.com/alitvinenko/ecareer_bot/internal/repository"
	"github.com/alitvinenko/ecareer_bot/internal/repository/club_member/converter"
	repoModel "github.com/alitvinenko/ecareer_bot/internal/repository/club_member/model"
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

func (r *repository) Get(ctx context.Context, ID int) (*model.ClubMember, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	var clubMember repoModel.ClubMember
	result := r.db.First(&clubMember, "id = ?", ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return converter.ToClubMemberFromRepo(&clubMember), nil
}

func (r *repository) GetByUsername(ctx context.Context, username string) (*model.ClubMember, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	var clubMember repoModel.ClubMember
	result := r.db.First(&clubMember, "username = ?", username)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return converter.ToClubMemberFromRepo(&clubMember), nil
}

func (r *repository) AddIfNotExists(ctx context.Context, clubMember *model.ClubMember) error {
	cm := converter.ToRepoFromClubMember(clubMember)

	result := r.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&cm)

	return result.Error
}
