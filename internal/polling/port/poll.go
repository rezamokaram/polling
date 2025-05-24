package port

import (
	"context"

	"polling/internal/polling/domain"
)

type Repo interface {
	CreatePoll(ctx context.Context, req domain.Poll) error
	PollList(ctx context.Context, req domain.Filter) ([]domain.Poll, error)
}
