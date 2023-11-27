package commands

import (
	"github.com/alitvinenko/ecareer_bot/internal/telegram/menu/buttons"
	tele "gopkg.in/telebot.v3"
)

func FirstLessonsHandler(c tele.Context) error {
	const message = `Уроки, с которых начинается Наставничество:
     
1️⃣- стартовый пакет новичка, все что нужно сделать до старта наставничества
2️⃣- обзор модулей
3️⃣- учебная программа
4️⃣- постановка намерения

Если что-то непонятно, напиши своему куратору.`

	if c.Callback() != nil {
		_ = c.Respond()
	}

	selector := &tele.ReplyMarkup{}
	selector.Inline(
		selector.Row(selector.URL("Стартовый пакет", "https://docs.google.com/presentation/d/1Kpz44Jf5MTgjg81MHEFr-EOFr2Nbz46nMGYPAYrAqf0/edit#slide=id.p")),
		selector.Row(selector.URL("Обзор модулей", "https://docs.google.com/presentation/d/1yPzTeyWMrA9_bRwaMqMh63UigZdw_6wbmgz4jr5vQS8/edit#slide=id.p")),
		selector.Row(selector.URL("Учебная программа", "https://docs.google.com/presentation/d/1nICFMCKXhaJGi7S-Al6hQLIKZ-_TiZvKQTN52gGevKg/edit#slide=id.g28ddb9c4fee_0_5")),
		selector.Row(selector.URL("Постановка намерения на наставничество", "https://docs.google.com/presentation/d/1sit4t1eYaG5Ushy8eqD_Royi3xhCjBMBasKX5dLdysk/edit#slide=id.p")),
		selector.Row(buttons.BackToStartBtn),
	)

	return c.Edit(message, &tele.SendOptions{
		ParseMode:   tele.ModeMarkdown,
		ReplyMarkup: selector,
	})
}
