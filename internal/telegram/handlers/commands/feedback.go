package commands

import (
	"github.com/alitvinenko/ecareer_bot/internal/telegram/menu/buttons"
	tele "gopkg.in/telebot.v3"
)

func FeedbackHandler(c tele.Context) error {
	const message = `Если у тебя есть идеи по улучшению моей работы и/или процесса онбординга - ты можешь написать их в топике [Чат-бот онбординга](https://t.me/c/1969859487/5708).

Моей разработкой занимается [Андрей](https://t.me/alitvinenko) @alitvinenko, продакт проекта [Наталия](https://t.me/hanerdy) @hanerdy.

Будем рады твоим идеям и обратной связи 🤗`

	if c.Callback() != nil {
		_ = c.Respond()
	}

	selector := &tele.ReplyMarkup{}
	selector.Inline(
		selector.Row(buttons.BackToStartBtn),
	)

	return c.Edit(message, &tele.SendOptions{
		ParseMode:   tele.ModeMarkdown,
		ReplyMarkup: selector,
	})
}
