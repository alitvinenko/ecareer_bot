package commands

import (
	"context"
	"github.com/alitvinenko/ecareer_bot/internal/service"
	"github.com/alitvinenko/ecareer_bot/internal/telegram/menu/buttons"
	tele "gopkg.in/telebot.v3"
	"log"
)

type StartCommandHandler struct {
	clubMemberService service.ClubMemberService
}

func NewStartCommandHandler(service service.ClubMemberService) *StartCommandHandler {
	return &StartCommandHandler{clubMemberService: service}
}

func (h *StartCommandHandler) Handle(c tele.Context) error {
	if c.Sender().IsBot {
		return nil
	}
	if !c.Message().Private() {
		return nil
	}

	// HACK
	userID := int(c.Sender().ID)
	username := c.Sender().Username

	_, err := h.clubMemberService.RegisterNewMember(context.Background(), userID, username)
	if err != nil {
		log.Printf("error on register new club member: %v", err)

		return c.Reply("При добавлении нового члена клуба произошла ошибка")
	}

	const message = `Привет, давай начнем работу.
Чем ты хочешь заняться?

*Твоя визитка* - заполним информацию о тебе, о твоем опыте и целях.
*Визитки участников* - для просмотра визитки участника клуба напиши /profile @юзернейм пользователя, я покажу всю информацию, которая у меня есть.
*Список вводных уроков* - ссылки на первые уроки наставничества чтобы начать погружаться
*Навигатор Клуба* - полезные ссылки на базу знаний, опросы, списки единомышленников и остальное интересное и полезное`

	selector := &tele.ReplyMarkup{}
	selector.Inline(
		selector.Row(buttons.AddProfileBtn),
		selector.Row(buttons.ProfileBtn),
		selector.Row(buttons.FirstLessonsBtn),
		selector.Row(buttons.ClubNavigateBtn),
		selector.Row(buttons.FeedbackBtn),
	)

	if c.Callback() != nil {
		_ = c.Respond()

		return c.Edit(message, &tele.SendOptions{
			ParseMode:   tele.ModeMarkdown,
			ReplyMarkup: selector,
		})
	}

	_ = c.Send(message, &tele.SendOptions{
		ParseMode:   tele.ModeMarkdown,
		ReplyMarkup: selector,
	})

	return nil
}
