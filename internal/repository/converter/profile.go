package converter

import (
	"github.com/alitvinenko/ecareer_bot/internal/model"
	repoModel "github.com/alitvinenko/ecareer_bot/internal/repository/club_member/model"
)

func ToProfileFromRepo(profile *repoModel.Profile) *model.Profile {
	return &model.Profile{
		ID:               profile.ID,
		WaitingForAnswer: profile.WaitingForAnswer,
		Data:             profile.Data,
		ClubMemberID:     profile.ClubMemberID,
	}
}

func ToRepoFromProfile(profile *model.Profile) *repoModel.Profile {
	return &repoModel.Profile{
		ID:               profile.ID,
		WaitingForAnswer: profile.WaitingForAnswer,
		Data:             profile.Data,
		ClubMemberID:     profile.ClubMemberID,
	}
}
