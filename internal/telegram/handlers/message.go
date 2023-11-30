package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/alitvinenko/ecareer_bot/internal/service"
	"github.com/alitvinenko/ecareer_bot/internal/telegram/menu/buttons"
	tele "gopkg.in/telebot.v3"
	"log"
)

type MessageHandler struct {
	clubMemberService service.ClubMemberService
}

func NewMessageHandler(clubMemberService service.ClubMemberService) *MessageHandler {
	return &MessageHandler{clubMemberService: clubMemberService}
}

func (h *MessageHandler) Handle(c tele.Context) error {
	if c.Message() == nil || c.Message().Text == "" {
		return nil
	}
	if !c.Message().Private() {
		return c.Send("Заполнить анкету ты можешь только в приватном чате со мной")
	}

	userID := int(c.Sender().ID)
	data := c.Message().Text

	err := h.clubMemberService.UpdateProfile(context.Background(), userID, data)
	if err != nil {
		log.Printf("error on save profile data: %v", err)

		if errors.Is(err, service.NotWaitingForAnswerErr) {
			return nil
		}

		_ = c.Reply("При сохранении анкеты произошел сбой. Повтори, пожалуйста, еще раз.")

		return nil
	}

	msgTemplate := "Спасибо! Твоя анкета сохранена.\n" +
		"\n" +
		"Теперь каждый участник группы может просмотреть ее выполнив команду `/profile %s`.\n" +
		"\n" +
		"*Содержимое твоей анкеты:*\n" +
		"\n" +
		"```" +
		"%s" +
		"```"
	msg := fmt.Sprintf(msgTemplate, c.Sender().Username, data)

	selector := &tele.ReplyMarkup{}
	selector.Inline(
		selector.Row(buttons.AddProfileConfirmBtn),
		selector.Row(buttons.BackToStartBtn),
	)

	err = c.Reply(msg, &tele.SendOptions{
		ParseMode:   tele.ModeMarkdown,
		ReplyMarkup: selector,
	})
	if err != nil {
		log.Printf("error on reply: %v", err)
	}

	return nil
}
