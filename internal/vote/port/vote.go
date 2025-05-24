package port

import (
	"context"
	"polling/internal/vote/domain"
)

type Repo interface {
	VotePoll(ctx context.Context, req domain.Vote) error
}
