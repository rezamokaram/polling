package storage

import (
	"context"
	"polling/internal/polling/domain"
	pollingPort "polling/internal/polling/port"
	"polling/pkg/adapters/storage/mapper"

	"gorm.io/gorm"
)

type pollRepo struct {
	db *gorm.DB
}

func NewPollRepo(db *gorm.DB) pollingPort.Repo {
	return &pollRepo{
		db: db,
	}
}

func (p *pollRepo) CreatePoll(ctx context.Context, req domain.Poll) error {
	poll := mapper.PollDomain2Storage(req)
	if err := p.db.Create(&poll).Error; err != nil {
		return err
	}

	return nil
}
