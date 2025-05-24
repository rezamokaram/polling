package mapper

import (
	"polling/internal/vote/domain"
	"polling/pkg/adapters/storage/types"
)

func DomainVote2Storage(vote domain.Vote, opt types.Option) types.Vote {
	return types.Vote{
		UserID:   vote.UserID,
		OptionID: opt.ID,
		PollID:   vote.PollID,
	}
}

func StorageVote2Domain(vote types.Vote) domain.Vote {
	return domain.Vote{
		PollID: vote.PollID,
		UserID: vote.UserID,
		Index:  0,
	}
}
