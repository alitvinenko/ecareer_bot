package commands

import (
	"context"
	"github.com/alitvinenko/ecareer_bot/internal/service"
	"github.com/alitvinenko/ecareer_bot/internal/telegram/menu/buttons"
	tele "gopkg.in/telebot.v3"
	"log"
)

type cancelAddProfileCommandHandler struct {
	profiles service.ProfileService
}

func NewCancelAddProfileCommandHandler(profiles service.ProfileService) *cancelAddProfileCommandHandler {
	return &cancelAddProfileCommandHandler{profiles: profiles}
}

func (h *cancelAddProfileCommandHandler) Handle(c tele.Context) error {
	ctx := context.Background()
	userID := c.Sender().ID

	err := h.profiles.StopWaitingProfileData(ctx, int(userID))
	if err != nil {
		log.Printf("error on stop waiting: %v", err)
	}

	selector := &tele.ReplyMarkup{}
	selector.Inline(
		selector.Row(buttons.BackToStartBtn),
	)

	return c.Send("Заполнение анкеты отменено.", &tele.SendOptions{
		ReplyMarkup: selector,
	})
}
