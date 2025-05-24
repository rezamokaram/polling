package http

import (
	"errors"
	"polling/api/pb"
	"polling/api/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func VotePoll(svcGetter ServiceGetter[*service.VoteService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		var req pb.VotePollRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		pollIdStr := c.Params("poll")
		pollId, err := strconv.ParseUint(pollIdStr, 10, 0)
		if err != nil {
			return fiber.ErrBadRequest
		}

		req.PollId = uint32(pollId)

		err = svc.VotePoll(c.UserContext(), &req)
		if err != nil {
			if errors.Is(err, service.ErrPollNotFound) {
				return c.SendStatus(fiber.StatusNotFound)
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.SendStatus(fiber.StatusCreated)
	}
}

func SkipPoll(svcGetter ServiceGetter[*service.VoteService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		var req pb.VotePollRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		pollIdStr := c.Params("poll")
		pollId, err := strconv.ParseUint(pollIdStr, 10, 0)
		if err != nil {
			return fiber.ErrBadRequest
		}

		req.PollId = uint32(pollId)
		req.OptionIndex = 0 // assume that 0 is the index for "skip" option

		err = svc.VotePoll(c.UserContext(), &req)
		if err != nil {
			if errors.Is(err, service.ErrPollNotFound) {
				return c.SendStatus(fiber.StatusNotFound)
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.SendStatus(fiber.StatusCreated)
	}
}
