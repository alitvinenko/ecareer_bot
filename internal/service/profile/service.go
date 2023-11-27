package profile

import (
	"context"
	"fmt"
	"github.com/alitvinenko/ecareer_bot/internal/model"
	"github.com/alitvinenko/ecareer_bot/internal/repository"
	def "github.com/alitvinenko/ecareer_bot/internal/service"
	"strings"
)

var _ def.ProfileService = (*service)(nil)

var NotWaitingForAnswerErr = fmt.Errorf("not waiting for answer")

type service struct {
	clubMembers repository.ClubMemberRepository
	profiles    repository.ProfileRepository
}

func NewService(clubMemberRepo repository.ClubMemberRepository, profileRepo repository.ProfileRepository) *service {
	return &service{clubMembers: clubMemberRepo, profiles: profileRepo}
}

func (s *service) FindByUsername(ctx context.Context, username string) (*model.Profile, error) {
	clubMember, err := s.clubMembers.GetByUsername(ctx, strings.Trim(username, "@"))
	if err != nil {
		return nil, err
	}
	if clubMember == nil {
		return nil, def.ClubMemberNotFoundErr
	}

	profile, err := s.profiles.Get(ctx, clubMember.ID)
	if err != nil {
		return nil, err
	}
	if profile == nil || profile.Data == "" {
		return nil, def.EmptyProfileServiceErr
	}

	return profile, nil
}

func (s *service) GetOrCreate(ctx context.Context, ID int) (*model.Profile, error) {
	profile, err := s.profiles.Get(ctx, ID)
	if err != nil {
		return nil, err
	}
	if profile == nil {
		profile = &model.Profile{
			ID: ID,
		}
		if err = s.profiles.Save(ctx, profile); err != nil {
			return nil, fmt.Errorf("error on create profile %d: %v", ID, err)
		}

		return profile, nil
	}

	return profile, nil
}

func (s *service) StartWaitingProfileData(ctx context.Context, profile *model.Profile) error {
	profile.WaitingForAnswer = true

	if err := s.profiles.Save(ctx, profile); err != nil {
		return fmt.Errorf("error on set waiting for anser = true: %v", err)
	}

	return nil
}

func (s *service) StopWaitingProfileData(ctx context.Context, userID int) error {
	profile, err := s.profiles.Get(ctx, userID)
	if err != nil {
		return fmt.Errorf("error on load profile %d: %v", userID, err)
	}
	if profile == nil {
		return nil
	}

	profile.WaitingForAnswer = false

	if err = s.profiles.Save(ctx, profile); err != nil {
		return fmt.Errorf("error on set waiting for anser = false: %v", err)
	}

	return nil
}

func (s *service) SaveProfileData(ctx context.Context, userID int, data string) error {
	profile, err := s.profiles.Get(ctx, userID)
	if err != nil {
		return fmt.Errorf("error on load profile %d: %v", userID, err)
	}
	if profile == nil {
		profile = &model.Profile{
			ID: userID,
		}
	}

	if !profile.WaitingForAnswer {
		return NotWaitingForAnswerErr
	}

	profile.Data = data
	profile.WaitingForAnswer = false

	if err = s.profiles.Save(ctx, profile); err != nil {
		return fmt.Errorf("error on update profile %d: %v", userID, err)
	}

	return nil
}
