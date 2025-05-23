package app

import (
	"context"
	"polling/config"
	pollingPort "polling/internal/polling/port"

	"gorm.io/gorm"
)

type App interface {
	PollingService(ctx context.Context) pollingPort.Service
	DB() *gorm.DB
	Config() config.PollingConfig
}
