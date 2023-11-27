package service

import (
	"context"
	"errors"
	"github.com/alitvinenko/ecareer_bot/internal/model"
)

var EmptyProfileServiceErr = errors.New("profile is empty")
var ClubMemberNotFoundErr = errors.New("club member not found")

type ClubMemberService interface {
	AddIfNotExists(ctx context.Context, id int, username string) error
}

type ProfileService interface {
	GetOrCreate(ctx context.Context, ID int) (*model.Profile, error)
	FindByUsername(ctx context.Context, username string) (*model.Profile, error)
	StartWaitingProfileData(ctx context.Context, profile *model.Profile) error
	StopWaitingProfileData(ctx context.Context, userID int) error
	SaveProfileData(ctx context.Context, userID int, data string) error
}
