package vote

import (
	"context"
	"polling/internal/vote/domain"
	"polling/internal/vote/port"
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) VotePoll(ctx context.Context, req domain.Vote) error {
	if err := s.repo.VotePoll(ctx, req); err != nil {
		return err
	}
	return nil
}
