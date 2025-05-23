package http

import (
	"context"
	"polling/api/service"
	"polling/app"
	"polling/config"
)

type ServiceGetter[T any] func(context.Context) T

func pollingServiceGetter(appContainer app.App, cfg config.POLLING) ServiceGetter[*service.PollingService] {
	return func(ctx context.Context) *service.PollingService {
		return service.NewPollingService(appContainer.PollingService(ctx))
	}
}
