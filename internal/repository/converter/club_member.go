package converter

import "github.com/alitvinenko/ecareer_bot/internal/model"
import repoModel "github.com/alitvinenko/ecareer_bot/internal/repository/club_member/model"

func ToClubMemberFromRepo(member *repoModel.ClubMember) *model.ClubMember {
	return &model.ClubMember{
		ID:       member.ID,
		Username: member.Username,
		Profile:  ToProfileFromRepo(member.Profile),
	}
}

func ToRepoFromClubMember(member *model.ClubMember) *repoModel.ClubMember {
	return &repoModel.ClubMember{
		ID:       member.ID,
		Username: member.Username,
		Profile:  ToRepoFromProfile(member.Profile),
	}
}
