package app

import (
	"github.com/alitvinenko/ecareer_bot/internal/config"
	"github.com/alitvinenko/ecareer_bot/internal/database"
	"github.com/alitvinenko/ecareer_bot/internal/repository"
	"github.com/alitvinenko/ecareer_bot/internal/repository/club_member"
	"github.com/alitvinenko/ecareer_bot/internal/service"
	club_member2 "github.com/alitvinenko/ecareer_bot/internal/service/club_member"
	tele "gopkg.in/telebot.v3"
	"gorm.io/gorm"
	"log"
	"time"
)

type serviceProvider struct {
	config *config.AppConfig

	tgBot *tele.Bot

	db *gorm.DB

	clubMemberRepository repository.ClubMemberRepository

	clubMemberService service.ClubMemberService
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) getConfig() *config.AppConfig {
	if s.config != nil {
		return s.config
	}

	appConfig, err := config.NewAppConfig()
	if err != nil {
		log.Fatalf("error on load appConfig: %s", err)
	}

	s.config = appConfig

	return s.config
}

func (s *serviceProvider) getTgBot() *tele.Bot {
	if s.tgBot != nil {
		return s.tgBot
	}

	pref := tele.Settings{
		Token: s.getConfig().GetToken(),
		Poller: &tele.LongPoller{
			Timeout:        10 * time.Second,
			AllowedUpdates: []string{"message", "poll", "poll_answer", "callback_query"},
		},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	s.tgBot = b

	return s.tgBot
}

func (s *serviceProvider) getDB() *gorm.DB {
	if s.db != nil {
		return s.db
	}

	s.db = database.Init(s.getConfig().GetDatabasePath())

	return s.db
}

func (s serviceProvider) getClubMemberRepository() repository.ClubMemberRepository {
	if s.clubMemberRepository != nil {
		return s.clubMemberRepository
	}

	s.clubMemberRepository = club_member.NewRepository(s.getDB())

	return s.clubMemberRepository
}

func (s *serviceProvider) getClubMemberService() service.ClubMemberService {
	if s.clubMemberService != nil {
		return s.clubMemberService
	}

	s.clubMemberService = club_member2.NewService(s.getClubMemberRepository())

	return s.clubMemberService
}
