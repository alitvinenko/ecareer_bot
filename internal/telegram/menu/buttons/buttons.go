package buttons

import "gopkg.in/telebot.v3"

const (
	startUnique             = "start"
	addProfileUnique        = "addprofile"
	addProfileConfirmUnique = "addprofile_confirm"
	profileUnique           = "profile"
	firstLessonsUnique      = "firstlessons"
	clubNavigateUnique      = "club_navigate"
	feedbackUnique          = "feedback"
)

var (
	StartBtn = telebot.Btn{
		Text:   "В меню",
		Unique: startUnique,
	}
	BackToStartBtn = telebot.Btn{
		Text:   "<< Вернуться в меню",
		Unique: startUnique,
	}
	AddProfileBtn = telebot.Btn{
		Text:   "📣 Твоя визитка",
		Unique: addProfileUnique,
	}
	AddProfileConfirmBtn = telebot.Btn{
		Text:   "Хочу прислать данные анкеты",
		Unique: "addprofile_yes",
	}
	ProfileBtn = telebot.Btn{
		Text:   "👤 Визитки участников",
		Unique: profileUnique,
	}
	FirstLessonsBtn = telebot.Btn{
		Text:   "Список вводных уроков",
		Unique: firstLessonsUnique,
	}
	ClubNavigateBtn = telebot.Btn{
		Text:   "Навигатор клуба",
		Unique: clubNavigateUnique,
	}
	BackToClubNavigateBtn = telebot.Btn{
		Text:   "< Вернуться в навигатор",
		Unique: clubNavigateUnique,
	}
	FeedbackBtn = telebot.Btn{
		Text:   "Пожелания",
		Unique: feedbackUnique,
	}
)
