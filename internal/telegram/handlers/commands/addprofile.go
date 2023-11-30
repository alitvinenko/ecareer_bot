package commands

import (
	"context"
	"fmt"
	"github.com/alitvinenko/ecareer_bot/internal/service"
	"github.com/alitvinenko/ecareer_bot/internal/telegram/menu/buttons"
	tele "gopkg.in/telebot.v3"
	"log"
)

type addProfileCommandHandler struct {
	clubMemberService service.ClubMemberService
}

func NewAddProfileCommandHandler(clubMemberService service.ClubMemberService) *addProfileCommandHandler {
	return &addProfileCommandHandler{clubMemberService: clubMemberService}
}

func (h *addProfileCommandHandler) Handle(c tele.Context) error {
	if c.Message().Chat.Type != tele.ChatPrivate {
		return c.Send("Заполнить анкету ты можешь только в приватном чате со мной")
	}

	ctx := context.Background()

	clubMember, err := h.clubMemberService.FindMemberByUsername(ctx, c.Sender().Username)
	if err != nil {
		log.Printf("error on load profile: %v", err)
		return c.Send("Ошибка на сервере")
	}

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

Для отмены заполнения выполни команду /cancel`

	o := &tele.SendOptions{
		ParseMode: tele.ModeMarkdown,
	}
	if clubMember.Profile.Empty() {
		selector := &tele.ReplyMarkup{}
		selector.Inline(
			selector.Row(buttons.AddProfileConfirmBtn),
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

	if !clubMember.Profile.Empty() {
		message := fmt.Sprintf("У меня уже есть твоя анкета. Ты хочешь ее заполнить снова?\n\n%s", clubMember.Profile.Data)

		selector := &tele.ReplyMarkup{}
		selector.Inline(
			selector.Row(buttons.AddProfileConfirmBtn),
			selector.Row(buttons.BackToStartBtn),
		)

		return c.Send(message, &tele.SendOptions{
			ParseMode:   tele.ModeMarkdown,
			ReplyMarkup: selector,
		})
	}

	return nil
}

func (h *addProfileCommandHandler) AddProfileConfirmHandle(c tele.Context) error {
	ctx := context.Background()

	clubMember, err := h.clubMemberService.FindMemberByUsername(ctx, c.Sender().Username)
	if err != nil {
		log.Printf("error on load profile: %v", err)
		return c.Send("Ошибка на сервере")
	}

	err = h.clubMemberService.StartWaitingProfileData(ctx, clubMember.ID)
	if err != nil {
		log.Printf("error on start waiting: %v", err)
		return c.Send("Ошибка на сервере")
	}

	message := `Отлично! Теперь пришли мне данные своей анкеты.

Если ты передумал и не хочешь заполнять анкету, выполни команду /cancel. В противном случае любое твое сообщение я буду считать данными анкеты.`
	return c.Reply(message, &tele.SendOptions{
		ParseMode: tele.ModeMarkdown,
	})
}
