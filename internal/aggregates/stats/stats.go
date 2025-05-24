package stats

import (
	"context"
	"polling/internal/polling/domain"
)

type AggregatePollStatsRepo interface {
	GetPollStats(ctx context.Context, pollID domain.PollID) (AggregatePollStats, error)
}

type AggregatePollStatsService interface {
	GetPollStats(ctx context.Context, pollID domain.PollID) (AggregatePollStats, error)
}

type AggregatePollStats struct {
	PollID     uint
	VotesStats []AggregateVoteStat
}

type AggregateVoteStat struct {
	OptionTitle string
	VotesCount  uint
}

type service struct {
	repo AggregatePollStatsRepo
}

func NewService(repo AggregatePollStatsRepo) AggregatePollStatsService {
	return &service{
		repo: repo,
	}
}

// IMPL

func (s *service) GetPollStats(ctx context.Context, pollID domain.PollID) (AggregatePollStats, error) {
	stats, err := s.repo.GetPollStats(ctx, pollID)
	if err != nil {
		return AggregatePollStats{}, err
	}
	return stats, nil
}
