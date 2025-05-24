package mapper

import (
	"polling/internal/polling/domain"
	"polling/pkg/adapters/storage/types"
)

func DomainOption2Storage(opt domain.Option, index uint) types.Option {
	return types.Option{
		Title:  opt.Title,
		PollID: uint(opt.PollID),
		Index:  index,
	}
}

func StorageOption2Domain(opt types.Option) domain.Option {
	return domain.Option{
		ID:     domain.OptionId(opt.ID),
		Title:  opt.Title,
		PollID: domain.PollID(opt.PollID),
		Index:  opt.Index,
	}
}
