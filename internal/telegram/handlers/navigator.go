package handlers

import (
	"github.com/alitvinenko/ecareer_bot/internal/telegram/menu/buttons"
	tele "gopkg.in/telebot.v3"
)

var (
	membersBtn = tele.Btn{
		Text:   "Участники клуба",
		Unique: "club_members",
	}
	themesBtn = tele.Btn{
		Text:   "Темы клуба",
		Unique: "club_themes",
	}
	pollsBtn = tele.Btn{
		Text:   "Опросы",
		Unique: "polls",
	}
	strategyBtn = tele.Btn{
		Text:   "Стратегия клуба",
		Unique: "club_strategy",
	}
)

type navigator struct {
}

func NewNavigator() *navigator {
	return &navigator{}
}

func (n *navigator) Handle(c tele.Context) error {
	const message = `Ты находишься в Навигаторе по Клубу. 

Здесь ты можешь познакомиться с другими участниками по кнопке "участники клуба" - клубы по городам, дни рождения и фото. 

Узнать какие темы обсуждаются в клубе по кнопке "темы клуба", например, где обсуждаются профессиональные темы, а где можно обсудить впечатления от уроков. 

По кнопке "опросы" ты получишь ссылки на несколько опросов, они не срочные, проходи когда буедт время и желание. 

Последняя кнопка, стратегия клуба" самая важная, она содердит ссылки на концепцию и стратегию клуба, как он создавался и как будет развиваться.`

	menu := &tele.ReplyMarkup{}
	menu.Inline(
		menu.Row(membersBtn),
		menu.Row(themesBtn),
		menu.Row(pollsBtn),
		menu.Row(strategyBtn),
		menu.Row(buttons.BackToStartBtn),
	)

	c.Bot().Handle(&membersBtn, n.membersBtnHandle)
	c.Bot().Handle(&themesBtn, n.themesBtnHandle)
	c.Bot().Handle(&pollsBtn, n.pollsBtnHandle)
	c.Bot().Handle(&strategyBtn, n.strategyBtnHandle)

	if c.Callback() != nil {
		_ = c.Respond()
	}

	_ = c.Edit(message, &tele.SendOptions{
		ParseMode:   tele.ModeMarkdown,
		ReplyMarkup: menu,
	})

	return nil
}

func (n *navigator) membersBtnHandle(c tele.Context) error {
	const message = `Здесь ты можешь найти своих земляков в клубах по городам, устроить оффлайн встречу или спросить куда можно сходить в этом городе.
Посмотреть у кого ДР в тот же день что и у тебя.

Увидеть фото других участников - они все размещены по тегу #фото`

	if c.Callback() != nil {
		_ = c.Respond()
	}

	selector := &tele.ReplyMarkup{}
	selector.Inline(
		selector.Row(selector.URL("Участники по городам", "https://t.me/c/1969859487/1/3736")),
		selector.Row(selector.URL("Дни рождения", "https://t.me/c/1969859487/1/333")),
		selector.Row(buttons.BackToClubNavigateBtn),
		selector.Row(buttons.BackToStartBtn),
	)

	return c.Edit(message, &tele.SendOptions{
		ParseMode:   tele.ModeMarkdown,
		ReplyMarkup: selector,
	})
}

func (n *navigator) themesBtnHandle(c tele.Context) error {
	const message = `Мини-комьюнити коллег по должности и сфере - CEO, COO,CTO, CPO, можно обсудить тренды и сложности, попросить совета.

Основное общение происходит в Болталке, здесь можно обсудить результаты тестов, обсудить уроки или вакансии

Как ориентироваться в топиках этого канала - см. здесь.

Прими участие в разработке фирменного мерча - фирменных наклеек и ежедневника в одноимённой ветке. (https://t.me/c/1969859487/2307)

База знаний - хранятся доп. уроки и практики от Катерины и партнёров Клуба и наша общая база рекомендуемых книг и курсов.`

	if c.Callback() != nil {
		_ = c.Respond()
	}

	selector := &tele.ReplyMarkup{}
	selector.Inline(
		selector.Row(selector.URL("Топики по должностям тут 👉", "https://t.me/c/1969859487/1/2791")),
		selector.Row(selector.URL("Болталка", "https://t.me/c/1969859487/2246")),
		selector.Row(selector.URL("База знаний", "https://t.me/c/1969859487/1866")),
		selector.Row(buttons.BackToClubNavigateBtn),
		selector.Row(buttons.BackToStartBtn),
	)

	return c.Edit(message, &tele.SendOptions{
		ParseMode:   tele.ModeMarkdown,
		ReplyMarkup: selector,
	})
}

func (n *navigator) pollsBtnHandle(c tele.Context) error {
	const message = `Поучаствуй в опросах

✔️Кто ты по Дизайну человека? (https://t.me/c/1969859487/1/955)
Построить свой бодиграф можно на сайте. (https://human-design.space/?ysclid=lh0kuyz3k2991203668)
✔️Какой ты Бог? (https://t.me/c/1969859487/1/949)
✔️Какая ты Богиня? (https://t.me/c/1969859487/1/951)
✔️Тест на уверенность в себе (https://t.me/c/1969859487/2246/6281) 

Это не срочно - как будет время и желание 😉`

	if c.Callback() != nil {
		_ = c.Respond()
	}

	selector := &tele.ReplyMarkup{}
	selector.Inline(
		selector.Row(buttons.BackToClubNavigateBtn),
		selector.Row(buttons.BackToStartBtn),
	)

	return c.Edit(message, &tele.SendOptions{
		ParseMode:   tele.ModeMarkdown,
		ReplyMarkup: selector,
	})
}

func (n *navigator) strategyBtnHandle(c tele.Context) error {
	const message = `Почитай и посмотри видео про концепцию и стратегию Клуба.

Как Катерина его создавала, для чего, что он даёт участникам, куда мы движемся - и многое другое в [Презентации концепции Клуба](https://drive.google.com/file/d/1UjkhC_hE0xMKQlqUBt99J3KOhUgQ6gOx/view).

И 2х часовой записи видео-презентации стратегии от 16 мая 2023: https://youtu.be/QMe3kHoUiJU

Если кто-то уже видел их в моём официальном канале - тут более полная версия.`

	if c.Callback() != nil {
		_ = c.Respond()
	}

	selector := &tele.ReplyMarkup{}
	selector.Inline(
		selector.Row(buttons.BackToClubNavigateBtn),
		selector.Row(buttons.BackToStartBtn),
	)

	return c.Edit(message, &tele.SendOptions{
		ParseMode:   tele.ModeMarkdown,
		ReplyMarkup: selector,
	})
}
