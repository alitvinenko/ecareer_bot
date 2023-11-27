package commands

import (
	"context"
	"fmt"
	"github.com/alitvinenko/ecareer_bot/internal/service"
	"github.com/alitvinenko/ecareer_bot/internal/telegram/menu/buttons"
	tele "gopkg.in/telebot.v3"
	"log"
)

var addProfileYesBtn = tele.Btn{
	Text:   "Хочу прислать данные анкеты",
	Unique: "addprofile_yes",
}

type addProfileCommandHandler struct {
	profiles service.ProfileService
}

func NewAddProfileCommandHandler(profiles service.ProfileService) *addProfileCommandHandler {
	return &addProfileCommandHandler{profiles: profiles}
}

func (h *addProfileCommandHandler) Handle(c tele.Context) error {
	if c.Message().Chat.Type != tele.ChatPrivate {
		return c.Send("Заполнить анкету вы можете только в приватном чате со мной")
	}

	ctx := context.Background()
	userID := int(c.Sender().ID)

	profile, err := h.profiles.GetOrCreate(ctx, userID)
	if err != nil {
		log.Printf("error on load profile: %v", err)
		return c.Send("Ошибка на сервере")
	}

	c.Bot().Handle(&addProfileYesBtn, h.processHandle)

	const message = `Заполни визитку о себе - и я представлю тебя другим резидентам 🥳 Ты можешь изменить визитку в любой момент прислав мне новый текст визитки по кнопке "Заполнить визитку".
Расскажу, кто с тобой из одного города, ниши, у кого такая же должность и карьерные или бизнес планы, есть ли у тебя тёзки и однофамильцы.

*Примеры визиток ты можешь посмотреть здесь - https://t.me/c/1969859487/1/337*
Тебе нужно подготовить текст о тебе, ответив на вопросы, ты можешь импровизировать или подготовиться заранее и изучить визитки других участников.

✔️Имя - фамилия, у нас есть повторяющиеся имена, поэтому фамилия обязательна 😉
✔️Город и в каких городах жил раньше 
✔️Должность текущая 
И цели - планы - мечты - к чему стремишься.
Например:
Project ➡️ Head of PMO ➡️ CEO
Если мечтаешь о своём бизнесе или уже есть - тоже пишите 🥳
✔️Кратко про свой карьерный путь - какие должности и компании есть в опыте.
✔️Ниши - это может быть полезно остальным.
Например: FinTech
✔️Ваши супер-силы, любимые скиллы и задачи.
✔️Хобби
✔️Ссылка на профиль Линкедин 😉

Для отмены заполнения напишите /cancel`

	o := &tele.SendOptions{
		ParseMode: tele.ModeMarkdown,
	}
	if profile.Data == "" {
		selector := &tele.ReplyMarkup{}
		selector.Inline(
			selector.Row(addProfileYesBtn),
			selector.Row(buttons.BackToStartBtn),
		)

		o.ReplyMarkup = selector
	}

	if c.Callback() != nil {
		_ = c.Respond()

		_ = c.Edit(message, o)
	} else {
		_ = c.Send(message, o)
	}

	if profile.Data != "" {
		message := fmt.Sprintf("У меня уже есть твоя анкета. Ты хочешь ее заполнить снова?\n\n%s", profile.Data)

		selector := &tele.ReplyMarkup{}
		selector.Inline(
			selector.Row(addProfileYesBtn),
			selector.Row(buttons.BackToStartBtn),
		)

		return c.Send(message, &tele.SendOptions{
			ParseMode:   tele.ModeMarkdown,
			ReplyMarkup: selector,
		})
	}

	return nil
}

func (h *addProfileCommandHandler) processHandle(c tele.Context) error {
	ctx := context.Background()
	userID := int(c.Sender().ID)

	profile, err := h.profiles.GetOrCreate(ctx, userID)
	if err != nil {
		log.Printf("error on load profile: %v", err)
		return c.Send("Ошибка на сервере")
	}

	err = h.profiles.StartWaitingProfileData(ctx, profile)
	if err != nil {
		log.Printf("error on start waiting: %v", err)
		return c.Send("Ошибка на сервере")
	}

	message := `Отлично! Теперь введите данные своей анкеты.

Если вы передумали и не хотите заполнять анкету, выполните команду /cancel. В противном случае, любое ваше сообщение я буду считать данными вашей анкеты.`
	return c.Reply(message, &tele.SendOptions{
		ParseMode: tele.ModeMarkdown,
	})
}
