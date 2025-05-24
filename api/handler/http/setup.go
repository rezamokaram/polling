package http

import (
	"fmt"
	"polling/app"
	"polling/config"

	"github.com/gofiber/fiber/v2"
)

func Run(appContainer app.App, cfg config.POLLING) error {
	router := fiber.New()

	api := router.Group("", setUserContext)

	pollingAPI(appContainer, cfg, api)
	voteAPI(appContainer, cfg, api)

	return router.Listen(fmt.Sprintf(":%d", cfg.Port))
}

func pollingAPI(appContainer app.App, cfg config.POLLING, router fiber.Router) {
	pollingSvcGetter := pollingServiceGetter(appContainer, cfg)
	router.Post("/polls", setTransaction(appContainer.DB()), CreatePoll(pollingSvcGetter))
	router.Get("/polls", setTransaction(appContainer.DB()), PollList(pollingSvcGetter))
}

func voteAPI(appContainer app.App, cfg config.POLLING, router fiber.Router) {
	voteSvcGetter := voteServiceGetter(appContainer, cfg)
	router.Post("/polls/:poll/vote", setTransaction(appContainer.DB()), VotePoll(voteSvcGetter))
	router.Post("/polls/:poll/skip", setTransaction(appContainer.DB()), SkipPoll(voteSvcGetter))
}
