package club_member

import (
	"context"
	"github.com/alitvinenko/ecareer_bot/internal/model"
	"github.com/alitvinenko/ecareer_bot/internal/repository"
)
import def "github.com/alitvinenko/ecareer_bot/internal/service"

var _ def.ClubMemberService = (*service)(nil)

type service struct {
	repository repository.ClubMemberRepository
}

func NewService(repository repository.ClubMemberRepository) *service {
	return &service{repository: repository}
}

func (s service) AddIfNotExists(ctx context.Context, id int, username string) error {
	return s.repository.AddIfNotExists(ctx, &model.ClubMember{
		ID:       id,
		Username: username,
	})
}
