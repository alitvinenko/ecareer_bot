package repository

import (
	"context"
	"github.com/alitvinenko/ecareer_bot/internal/model"
)

type ClubMemberRepository interface {
	Get(ctx context.Context, ID int) (*model.ClubMember, error)
	GetByUsername(ctx context.Context, username string) (*model.ClubMember, error)
	AddIfNotExists(ctx context.Context, clubMember *model.ClubMember) error
}

type ProfileRepository interface {
	Get(ctx context.Context, ID int) (*model.Profile, error)
	Save(ctx context.Context, profile *model.Profile) error
}
