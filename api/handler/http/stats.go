package http

import (
	"errors"
	"polling/api/pb"
	"polling/api/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func PollStats(svcGetter ServiceGetter[*service.StatsService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		var req pb.PollStatsRequest

		pollIdStr := c.Params("poll")
		pollId, err := strconv.ParseUint(pollIdStr, 10, 0)
		if err != nil {
			return fiber.ErrBadRequest
		}

		req.PollId = uint32(pollId)

		resp, err := svc.PollStats(c.UserContext(), &req)
		if err != nil {
			if errors.Is(err, service.ErrPollNotFound) {
				return c.SendStatus(fiber.StatusNotFound)
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(&resp)
	}
}
