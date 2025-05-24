package app

import (
	"context"
	"polling/config"
	pollingPort "polling/internal/polling/port"
	votePort "polling/internal/vote/port"

	"gorm.io/gorm"
)

type App interface {
	PollingService(ctx context.Context) pollingPort.Service
	VoteService(ctx context.Context) votePort.Service
	DB() *gorm.DB
	Config() config.PollingConfig
}
