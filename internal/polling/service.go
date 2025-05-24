package polling

import (
	"context"
	"errors"
	"polling/internal/polling/domain"
	"polling/internal/polling/port"
)

var (
	ErrPollNotFound = errors.New("poll not found")
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreatePoll(ctx context.Context, req domain.Poll) error {
	if err := s.repo.CreatePoll(ctx, req); err != nil {
		return err
	}
	return nil
}

func (s *service) PollList(ctx context.Context, req domain.Filter) ([]domain.Poll, error) {
	list, err := s.repo.PollList(ctx, req)
	if err != nil {
		return nil, err
	}
	return list, nil
}
