package storage

import (
	"context"
	"log"
	"polling/internal/polling/domain"
	pollingPort "polling/internal/polling/port"
	"polling/pkg/adapters/storage/mapper"
	"polling/pkg/adapters/storage/types"

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

func (p *pollRepo) PollList(ctx context.Context, req domain.Filter) ([]domain.Poll, error) {
	query := p.db.Model(&types.Poll{}).WithContext(ctx)
	if req.Tag.Title != "" {
		query = query.Joins("JOIN tags ON tags.poll_id = polls.id").
			Where("tags.title = ?", req.Tag.Title).Preload("Tags")
	}

	if req.Limit <= 0 {
		req.Limit = 1
	}
	query = query.Limit(int(req.Limit))

	if req.Page > 0 {
		query = query.Offset((int(req.Page) - 1) * int(req.Limit))
	}

	var list []types.Poll
	if err := query.Preload("Options").Find(&list).Error; err != nil {
		log.Println("err: ", err.Error())
		return nil, err
	}

	res := make([]domain.Poll, len(list))
	for i, poll := range list {
		res[i] = mapper.PollStorage2Domain(poll)
	}

	return res, nil
}
