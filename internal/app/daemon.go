package app

import (
	"context"
	"github.com/alitvinenko/ecareer_bot/internal/telegram/handlers"
	"github.com/alitvinenko/ecareer_bot/internal/telegram/handlers/commands"
	"github.com/alitvinenko/ecareer_bot/internal/telegram/menu/buttons"
	tele "gopkg.in/telebot.v3"
	"log"
)

type Daemon struct {
	serviceProvider *serviceProvider
}

func NewDaemon(ctx context.Context) *Daemon {
	app := &Daemon{}

	err := app.initServiceProvider(ctx)
	if err != nil {
		log.Fatalf("error on init service provider: %v", err)
	}
	err = app.setBotHandlers()
	if err != nil {
		log.Fatalf("error on init bot commands handlers: %v", err)
	}

	return app
}

func (d Daemon) Run(_ context.Context) error {
	log.Printf("Authorized on account %s", d.serviceProvider.getTgBot().Me.Username)

	d.serviceProvider.getTgBot().Start()

	return nil
}

func (a *Daemon) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *Daemon) setBotHandlers() error {

	b := a.serviceProvider.getTgBot()

	clubMemberService := a.serviceProvider.getClubMemberService()

	joinUserHandler := handlers.NewJoinUserHandler(clubMemberService)
	startCommandHandler := commands.NewStartCommandHandler(clubMemberService)
	addProfileHandler := commands.NewAddProfileCommandHandler(clubMemberService)
	cancelAddProfileHandler := commands.NewCancelAddProfileCommandHandler(clubMemberService)
	profileCommandHandler := commands.NewProfileCommandHandler(clubMemberService)
	clubNavigatorHandler := handlers.NewNavigator()
	messageHandler := handlers.NewMessageHandler(clubMemberService)

	b.Handle("/start", startCommandHandler.Handle)
	b.Handle("/addprofile", addProfileHandler.Handle)
	b.Handle("/cancel", cancelAddProfileHandler.Handle)
	b.Handle("/profile", profileCommandHandler.Handle)

	b.Handle(tele.OnText, messageHandler.Handle)
	b.Handle(tele.OnUserJoined, joinUserHandler.Handle)

	b.Handle(&buttons.StartBtn, startCommandHandler.Handle)
	b.Handle(&buttons.AddProfileBtn, addProfileHandler.Handle)
	b.Handle(&buttons.AddProfileConfirmBtn, addProfileHandler.AddProfileConfirmHandle)
	b.Handle(&buttons.FeedbackBtn, commands.FeedbackHandler)
	b.Handle(&buttons.FirstLessonsBtn, commands.FirstLessonsHandler)
	b.Handle(&buttons.ProfileBtn, profileCommandHandler.Handle)
	b.Handle(&buttons.ClubNavigateBtn, clubNavigatorHandler.Handle)

	return nil
}
