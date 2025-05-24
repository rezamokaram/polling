package port

import (
	"context"
	"polling/internal/vote/domain"
)

type Service interface {
	VotePoll(ctx context.Context, req domain.Vote) error
	// SkipPoll(ctx context.Context, req domain.Filter) ([]domain.Poll, error)
}
