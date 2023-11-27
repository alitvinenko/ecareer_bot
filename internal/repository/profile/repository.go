package profile

import (
	"context"
	"errors"
	"github.com/alitvinenko/ecareer_bot/internal/model"
	def "github.com/alitvinenko/ecareer_bot/internal/repository"
	"github.com/alitvinenko/ecareer_bot/internal/repository/profile/converter"
	repoModel "github.com/alitvinenko/ecareer_bot/internal/repository/profile/model"
	"gorm.io/gorm"
	"sync"
)

var _ def.ProfileRepository = (*repository)(nil)

type repository struct {
	db *gorm.DB
	m  sync.RWMutex
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Get(ctx context.Context, ID int) (*model.Profile, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	var profile repoModel.Profile
	result := r.db.Debug().First(&profile, "id = ?", ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return converter.ToProfileFromRepo(&profile), nil
}

func (r *repository) Save(ctx context.Context, profile *model.Profile) error {
	r.m.Lock()
	defer r.m.Unlock()

	p := converter.ToRepoFromProfile(profile)

	result := r.db.Debug().Save(&p)

	return result.Error
}
