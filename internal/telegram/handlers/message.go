package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/alitvinenko/ecareer_bot/internal/service"
	"github.com/alitvinenko/ecareer_bot/internal/service/profile"
	"github.com/alitvinenko/ecareer_bot/internal/telegram/menu/buttons"
	tele "gopkg.in/telebot.v3"
	"log"
)

type MessageHandler struct {
	profileService service.ProfileService
}

func NewMessageHandler(profileService service.ProfileService) *MessageHandler {
	return &MessageHandler{profileService: profileService}
}

func (h *MessageHandler) Handle(c tele.Context) error {
	if c.Message() == nil || c.Message().Text == "" {
		return nil
	}
	if c.Message().Chat.Type != tele.ChatPrivate {
		return c.Send("Заполнить анкету вы можете только в приватном чате со мной")
	}

	userID := int(c.Sender().ID)
	data := c.Message().Text

	err := h.profileService.SaveProfileData(context.Background(), userID, data)
	if err != nil {
		log.Printf("error on save profile data: %v", err)

		if errors.Is(err, profile.NotWaitingForAnswerErr) {
			return nil
		}

		_ = c.Reply("При сохранении анкеты произошел сбой. Повторите, пожалуйста, еще раз.")
	}

	message := `Спасибо! Ваша анкета сохранена. 

Теперь каждый участник группы может просмотреть вашу анкету выполнив команду /profile %s.

*Содержимое вашей анкеты:*

%s`

	selector := &tele.ReplyMarkup{}
	selector.Inline(
		selector.Row(buttons.BackToStartBtn),
	)

	_ = c.Reply(fmt.Sprintf(message, c.Sender().Username, data), &tele.SendOptions{
		ParseMode:   tele.ModeMarkdown,
		ReplyMarkup: selector,
	})

	return nil
}
