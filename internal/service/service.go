package service

import (
	"context"
	"errors"
	"github.com/alitvinenko/ecareer_bot/internal/model"
)

var EmptyProfileServiceErr = errors.New("profile is empty")
var ClubMemberNotFoundErr = errors.New("club member not found")
var NotWaitingForAnswerErr = errors.New("not waiting for answer")

type ClubMemberService interface {
	RegisterNewMember(ctx context.Context, id int, username string) (*model.ClubMember, error)
	FindMemberByUsername(ctx context.Context, username string) (*model.ClubMember, error)
	UpdateProfile(ctx context.Context, ID int, profileData string) (err error)
	StartWaitingProfileData(ctx context.Context, ID int) error
	StopWaitingProfileData(ctx context.Context, ID int) error
}
