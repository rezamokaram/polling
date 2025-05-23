package port

import (
	"context"
	"polling/internal/polling/domain"
)

type Service interface {
	CreatePoll(ctx context.Context, req domain.Poll) error
}
