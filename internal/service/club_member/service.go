package club_member

import (
	"context"
	"github.com/alitvinenko/ecareer_bot/internal/lib/e"
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

func (s service) RegisterNewMember(ctx context.Context, id int, username string) (*model.ClubMember, error) {
	clubMember, err := s.repository.Get(ctx, id)
	if err != nil {
		return nil, e.Wrap("error on load club member", err)
	}
	if clubMember != nil {
		return clubMember, nil
	}

	clubMember = &model.ClubMember{
		ID:       id,
		Username: username,
		Profile: &model.Profile{
			ID: id,
		},
	}

	if err = s.repository.Create(ctx, clubMember); err != nil {
		return nil, e.Wrap("error on save new club member", err)
	}

	return clubMember, nil
}

func (s *service) FindMemberByUsername(ctx context.Context, username string) (*model.ClubMember, error) {
	clubMember, err := s.repository.GetByUsername(ctx, username)
	if err != nil {
		return nil, e.Wrap("error on find club member by username", err)
	}

	return clubMember, nil
}

func (s *service) UpdateProfile(ctx context.Context, ID int, profileData string) (err error) {
	defer func() { err = e.WrapIfErr("error on update profile", err) }()

	clubMember, err := s.repository.Get(ctx, ID)
	if err != nil {
		return err
	}
	if clubMember.ID == 0 {
		return def.ClubMemberNotFoundErr
	}
	if !clubMember.Profile.WaitingForAnswer {
		return def.NotWaitingForAnswerErr
	}

	clubMember.Profile.Data = profileData
	clubMember.Profile.WaitingForAnswer = false

	if err = s.repository.Update(ctx, clubMember); err != nil {
		return err
	}

	return nil
}

func (s *service) StartWaitingProfileData(ctx context.Context, ID int) (err error) {
	defer func() { err = e.WrapIfErr("error on start waiting profile data", err) }()
	clubMember, err := s.repository.Get(ctx, ID)
	if err != nil {
		return err
	}
	if clubMember.ID == 0 {
		return def.ClubMemberNotFoundErr
	}

	clubMember.Profile.WaitingForAnswer = true
	err = s.repository.Update(ctx, clubMember)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) StopWaitingProfileData(ctx context.Context, ID int) (err error) {
	defer func() { err = e.WrapIfErr("error on start waiting profile data", err) }()
	clubMember, err := s.repository.Get(ctx, ID)
	if err != nil {
		return err
	}
	if clubMember.ID == 0 {
		return def.ClubMemberNotFoundErr
	}

	clubMember.Profile.WaitingForAnswer = false
	err = s.repository.Update(ctx, clubMember)
	if err != nil {
		return err
	}

	return nil
}
