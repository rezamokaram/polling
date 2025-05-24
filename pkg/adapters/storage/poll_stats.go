package storage

import (
	"context"
	"log"
	"polling/internal/aggregates/stats"
	"polling/internal/polling/domain"
	"polling/pkg/adapters/storage/types"

	"gorm.io/gorm"
)

type pollStatsRepo struct {
	db *gorm.DB
}

func NewPollStatsRepo(db *gorm.DB) stats.AggregatePollStatsRepo {
	return &pollStatsRepo{
		db: db,
	}
}

func (p *pollStatsRepo) GetPollStats(ctx context.Context, req domain.PollID) (stats.AggregatePollStats, error) {
	pollId := uint(req)

	var poll types.Poll
	if err := p.db.WithContext(ctx).Model(&types.Poll{}).Where("id = ?", pollId).Preload("Options").First(&poll).Error; err != nil {
		return stats.AggregatePollStats{}, err
	}

	resp := stats.AggregatePollStats{
		PollID:     pollId,
		VotesStats: make([]stats.AggregateVoteStat, len(poll.Options)),
	}
	for i, opt := range poll.Options {
		count := int64(0)
		if err := p.db.WithContext(ctx).Model(&types.Vote{}).Where("option_id = ?", opt.ID).Count(&count).Error; err != nil {
			log.Printf("error counting votes for option %d: %v", opt.ID, err)
		}

		resp.VotesStats[i] = stats.AggregateVoteStat{
			OptionTitle: opt.Title,
			VotesCount:  uint(count),
		}
	}
	// count := int64(0)
	// if err := p.db.WithContext(ctx).Model(&types.Vote{}).Where("poll_id = ? AND option_index = ?", pollId, 0).Count(&count).Error; err != nil {
	// 	log.Printf("error counting skipped votes for poll %d: %v", pollId, err)
	// }
	// resp.VotesStats[len(poll.Options)] = stats.AggregateVoteStat{
	// 	OptionTitle: "Skipped",
	// 	VotesCount:  uint(count),
	// }

	return resp, nil
}
