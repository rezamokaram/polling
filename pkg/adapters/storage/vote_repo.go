package storage

import (
	"context"
	"log"
	"polling/internal/vote/domain"
	"polling/internal/vote/port"
	"polling/pkg/adapters/storage/mapper"
	"polling/pkg/adapters/storage/types"

	"gorm.io/gorm"
)

type voteRepo struct {
	db *gorm.DB
}

func NewVoteRepo(db *gorm.DB) port.Repo {
	return &voteRepo{
		db: db,
	}
}

func (p *voteRepo) VotePoll(ctx context.Context, req domain.Vote) error {
	var opt types.Option
	if req.Index != 0 {
		if err := p.db.WithContext(ctx).Model(&types.Option{}).Where("index = ?", req.Index).First(&opt).Error; err != nil {
			log.Printf("error fetching option with index %d: %v", req.Index, err)
			return domain.ErrOptionNotFound
		}
	}

	var poll types.Poll
	if err := p.db.WithContext(ctx).Model(&types.Poll{}).Where("id = ?", req.PollID).First(&poll).Error; err != nil {
		log.Printf("error fetching poll with ID %d: %v", req.PollID, err)
		return domain.ErrPollNotFound
	}
	vote := mapper.DomainVote2Storage(req, opt)

	var existingVote types.Vote
	if err := p.db.WithContext(ctx).Where("user_id = ? AND poll_id = ?", req.UserID, req.PollID).First(&existingVote).Error; err == nil {
		return domain.ErrAlreadyVoted
	}

	if err := p.db.WithContext(ctx).Create(&vote).Error; err != nil {
		return err
	}

	return nil
}
