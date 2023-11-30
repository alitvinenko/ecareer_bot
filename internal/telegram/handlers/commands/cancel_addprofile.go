package commands

import (
	"context"
	"github.com/alitvinenko/ecareer_bot/internal/service"
	"github.com/alitvinenko/ecareer_bot/internal/telegram/menu/buttons"
	tele "gopkg.in/telebot.v3"
	"log"
)

type cancelAddProfileCommandHandler struct {
	clubMemberService service.ClubMemberService
}

func NewCancelAddProfileCommandHandler(clubMemberService service.ClubMemberService) *cancelAddProfileCommandHandler {
	return &cancelAddProfileCommandHandler{clubMemberService: clubMemberService}
}

func (h *cancelAddProfileCommandHandler) Handle(c tele.Context) error {
	ctx := context.Background()
	userID := c.Sender().ID

	err := h.clubMemberService.StopWaitingProfileData(ctx, int(userID))
	if err != nil {
		log.Printf("error on stop waiting: %v", err)

		return c.Send("Отменить заполнение анкеты не удалось. Обратитесь в поддержку.")
	}

	selector := &tele.ReplyMarkup{}
	selector.Inline(
		selector.Row(buttons.BackToStartBtn),
	)

	return c.Send("Заполнение анкеты отменено.", &tele.SendOptions{
		ReplyMarkup: selector,
	})
}
