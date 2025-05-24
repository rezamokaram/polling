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

	registerAuthAPI(appContainer, cfg, api)

	return router.Listen(fmt.Sprintf(":%d", cfg.Port))
}

func registerAuthAPI(appContainer app.App, cfg config.POLLING, router fiber.Router) {
	pollingSvcGetter := pollingServiceGetter(appContainer, cfg)
	router.Post("/polls", setTransaction(appContainer.DB()), CreatePoll(pollingSvcGetter))
	router.Get("/polls", setTransaction(appContainer.DB()), PollList(pollingSvcGetter))
}
