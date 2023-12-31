package commands

import (
	"context"
	"github.com/alitvinenko/ecareer_bot/internal/service"
	"github.com/alitvinenko/ecareer_bot/internal/telegram/menu/buttons"
	tele "gopkg.in/telebot.v3"
	"log"
	"strings"
)

type ProfileCommandHandler struct {
	clubMemberService service.ClubMemberService
}

func NewProfileCommandHandler(clubMemberService service.ClubMemberService) *ProfileCommandHandler {
	return &ProfileCommandHandler{clubMemberService: clubMemberService}
}

func (h *ProfileCommandHandler) Handle(c tele.Context) error {
	tags := c.Args()
	if len(tags) == 1 && tags[0] != "" {
		_ = h.showProfile(strings.Trim(tags[0], "@"), c)
	} else {
		_ = h.baseHandle(c)
	}

	return nil
}

func (h *ProfileCommandHandler) baseHandle(c tele.Context) error {
	const message = `Для просмотра визитки участника клуба напиши /profile @юзернейм пользователя.

Либо посмотри список всех участников с кликабельными ссылками на их подробные визитки [тут](https://t.me/c/1969859487/1/337), или в тематических топиках:
✔️[СРО, продукт](https://t.me/c/1969859487/2248/5621)
✔️[СТО, CIO, CDTO](https://t.me/c/1969859487/2247/5618)
✔️[СОО](https://t.me/c/1969859487/2249/5623)
✔️[РМО](https://t.me/c/1969859487/2251/5616)
✔️[Свой бизнес, консалтинг](https://t.me/c/1969859487/2252/9537)
✔️[Фаундеры ИТ стартапов](https://t.me/c/1969859487/2252/9538)`

	selector := &tele.ReplyMarkup{}
	selector.Inline(
		selector.Row(buttons.BackToStartBtn),
	)

	if c.Callback() != nil {
		_ = c.Respond()

		return c.Edit(message, &tele.SendOptions{
			ParseMode:   tele.ModeMarkdown,
			ReplyMarkup: selector,
		})
	}

	return c.Send(message, &tele.SendOptions{
		ParseMode:   tele.ModeMarkdown,
		ReplyMarkup: selector,
	})
}

func (h *ProfileCommandHandler) showProfile(username string, c tele.Context) error {
	clubMember, err := h.clubMemberService.FindMemberByUsername(context.Background(), username)
	if err != nil {
		log.Printf("error on load profile %s: %v", username, err)
		_ = c.Reply("На сервере произошла ошибка")

		return err
	}
	if clubMember == nil {
		message := `Не вижу такого пользователя в клубе, проверь написание имени и попробуй еще раз.`
		_ = c.Reply(message)

		return nil
	}

	if clubMember.Profile.Empty() {
		message := `Участник еще не заполнил визитку, у меня нет информации.`
		_ = c.Reply(message)

		return nil
	}

	return c.Reply(clubMember.Profile.Data, &tele.SendOptions{
		ParseMode: tele.ModeMarkdown,
	})
}
