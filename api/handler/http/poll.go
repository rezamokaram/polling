package http

import (
	"errors"
	"polling/api/pb"
	"polling/api/service"

	"github.com/gofiber/fiber/v2"
)

func CreatePoll(svcGetter ServiceGetter[*service.PollingService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		var req pb.CreatePollRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		err := svc.CreatePoll(c.UserContext(), &req)
		if err != nil {
			if errors.Is(err, service.ErrPollNotFound) {
				return c.SendStatus(fiber.StatusNotFound)
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.SendStatus(fiber.StatusCreated)
	}
}
