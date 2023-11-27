package handlers

import (
	"context"
	"fmt"
	"github.com/alitvinenko/ecareer_bot/internal/service"
	tele "gopkg.in/telebot.v3"
	"log"
)

type JoinUserHandler struct {
	clubMemberService service.ClubMemberService
}

func NewJoinUserHandler(clubMemberService service.ClubMemberService) *JoinUserHandler {
	return &JoinUserHandler{clubMemberService: clubMemberService}
}

func (h *JoinUserHandler) Handle(c tele.Context) error {
	if c.Message().UserJoined.IsBot {
		return nil
	}

	userID := int(c.Message().UserJoined.ID)
	username := c.Message().UserJoined.Username

	err := h.clubMemberService.AddIfNotExists(context.Background(), userID, username)
	if err != nil {
		log.Printf("error on load or create a new chat member: %v", err)

		return c.Reply("При добавлении нового члена клуба произошла ошибка")
	}

	message := `👋 Привет! 
    
Я *бот-Карьерист* - главный помощник и твой навигатор по Клубу «Карьера и Бизнес с Екатериной Евстафьевой» и Наставничеству «Карьера мечты». 
    
Я познакомлю тебя с участниками и традициями Клуба - и расскажу, как мы будем работать в наставничестве. 
    
Сейчас я могу помочь тебе:
1️⃣ рассказать подробно о стратегии и концепции Клуба
2️⃣ заполнить твою визитку
3️⃣ показать визитки участников клуба, показать фото
4️⃣ дать ссылки на вводные уроки в наставничестве
5️⃣ рассказать и показать где что лежит: база знаний, чаты единомышленников и прочее интересное, что можно найти в клубе.

Для этого нажми на кнопку "Начать"`

	selector := &tele.ReplyMarkup{}
	selector.Inline(
		selector.Row(selector.URL("Начать", fmt.Sprintf("tg://resolve?domain=%s", c.Bot().Me.Username))),
	)

	return c.Reply(message, &tele.SendOptions{
		ParseMode:   tele.ModeMarkdown,
		ReplyMarkup: selector,
	})
}
