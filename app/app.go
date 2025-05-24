package app

import (
	"context"
	"fmt"
	"polling/config"
	"polling/internal/polling"
	pollingPort "polling/internal/polling/port"
	"polling/internal/vote"
	votePort "polling/internal/vote/port"
	redisAdapter "polling/pkg/adapters/cache"
	"polling/pkg/adapters/storage"
	"polling/pkg/adapters/storage/types"
	"polling/pkg/cache"
	appCtx "polling/pkg/context"
	"polling/pkg/postgres"

	"gorm.io/gorm"
)

type app struct {
	db             *gorm.DB
	redisProvider  cache.Provider
	cfg            config.PollingConfig
	pollingService pollingPort.Service
	voteService    votePort.Service
}

func NewApp(cfg config.PollingConfig) (App, error) {
	a := &app{
		cfg: cfg,
	}

	if err := a.setDB(); err != nil {
		return nil, err
	}

	a.setRedis()

	return a, nil
}

func NewMustApp(cfg config.PollingConfig) App {
	app, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return app
}

func (a *app) setDB() error {
	db, err := postgres.NewPsqlGormConnection(postgres.DBConnOptions{
		User:   a.cfg.Postgres.User,
		Pass:   a.cfg.Postgres.Password,
		Host:   a.cfg.Postgres.Host,
		Port:   a.cfg.Postgres.Port,
		DBName: a.cfg.Postgres.DB,
		Schema: a.cfg.Postgres.Schema,
	})

	if err != nil {
		return err
	}

	if err := db.AutoMigrate(
		&types.Poll{},
		&types.Option{},
		&types.Tag{},
		&types.Vote{},
	); err != nil {
		return err
	}

	a.db = db
	return nil
}

func (a *app) setRedis() {
	a.redisProvider = redisAdapter.NewRedisProvider(fmt.Sprintf("%s:%d", a.cfg.Redis.Host, a.cfg.Redis.Port))
}

func (a *app) pollingServiceWithDB(db *gorm.DB) pollingPort.Service {
	return polling.NewService(storage.NewPollRepo(db))
}

func (a *app) voteServiceWithDB(db *gorm.DB) votePort.Service {
	return vote.NewService(storage.NewVoteRepo(db))
}

// IMPL

func (a *app) PollingService(ctx context.Context) pollingPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.pollingService == nil {
			a.pollingService = a.pollingServiceWithDB(a.db)
		}
		return a.pollingService
	}

	return a.pollingServiceWithDB(db)
}

func (a *app) VoteService(ctx context.Context) votePort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.pollingService == nil {
			a.pollingService = a.pollingServiceWithDB(a.db)
		}
		return a.voteService
	}

	return a.voteServiceWithDB(db)
}

func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) Config() config.PollingConfig {
	return a.cfg
}
