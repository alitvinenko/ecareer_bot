package buttons

import "gopkg.in/telebot.v3"

var (
	StartBtn = telebot.Btn{
		Text:   "В меню",
		Unique: "start",
	}
	BackToStartBtn = telebot.Btn{
		Text:   "<< Вернуться в меню",
		Unique: "start",
	}
	AddProfileBtn = telebot.Btn{
		Text:   "📣 Заполнить визитку",
		Unique: "addprofile",
	}
	ProfileBtn = telebot.Btn{
		Text:   "👤 Визитки участников",
		Unique: "profile",
	}
	FirstLessonsBtn = telebot.Btn{
		Text:   "Список вводных уроков",
		Unique: "firstlessons",
	}
	ClubNavigateBtn = telebot.Btn{
		Text:   "Навигатор клуба",
		Unique: "club_navigate",
	}
	BackToClubNavigateBtn = telebot.Btn{
		Text:   "< Вернуться в навигатор",
		Unique: "club_navigate",
	}
	FeedbackBtn = telebot.Btn{
		Text:   "Пожелания",
		Unique: "feedback",
	}
)
