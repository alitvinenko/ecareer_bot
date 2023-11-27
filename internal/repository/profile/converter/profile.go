package converter

import "github.com/alitvinenko/ecareer_bot/internal/model"
import repoModel "github.com/alitvinenko/ecareer_bot/internal/repository/profile/model"

func ToProfileFromRepo(profile *repoModel.Profile) *model.Profile {
	return &model.Profile{
		ID:               profile.ID,
		WaitingForAnswer: profile.WaitingForAnswer,
		Data:             profile.Data,
	}
}

func ToRepoFromProfile(profile *model.Profile) *repoModel.Profile {
	return &repoModel.Profile{
		ID:               profile.ID,
		WaitingForAnswer: profile.WaitingForAnswer,
		Data:             profile.Data,
	}
}
