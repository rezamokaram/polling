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

func voteServiceGetter(appContainer app.App, cfg config.POLLING) ServiceGetter[*service.VoteService] {
	return func(ctx context.Context) *service.VoteService {
		return service.NewVoteService(appContainer.VoteService(ctx))
	}
}

func statsServiceGetter(appContainer app.App, cfg config.POLLING) ServiceGetter[*service.StatsService] {
	return func(ctx context.Context) *service.StatsService {
		return service.NewStatsService(appContainer.StatsService(ctx))
	}
}
