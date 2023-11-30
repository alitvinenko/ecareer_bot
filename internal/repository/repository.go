package repository

import (
	"context"
	"github.com/alitvinenko/ecareer_bot/internal/model"
)

type ClubMemberRepository interface {
	Create(ctx context.Context, member *model.ClubMember) error
	Get(ctx context.Context, ID int) (*model.ClubMember, error)
	GetByUsername(ctx context.Context, username string) (*model.ClubMember, error)
	Update(ctx context.Context, member *model.ClubMember) error
}
