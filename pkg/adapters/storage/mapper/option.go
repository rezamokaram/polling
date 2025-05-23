package mapper

import (
	"polling/internal/polling/domain"
	"polling/pkg/adapters/storage/types"
)

func DomainOption2Storage(opt domain.Option) types.Option {
	return types.Option{
		Title: opt.Title,
	}
}

func StorageOption2Domain(opt types.Option) domain.Option {
	return domain.Option{
		ID:    domain.OptionId(opt.ID),
		Title: opt.Title,
	}
}
