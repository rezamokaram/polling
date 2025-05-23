package mapper

import (
	"polling/internal/polling/domain"
	"polling/pkg/adapters/storage/types"
)

func PollDomain2Storage(poll domain.Poll) types.Poll {
	options := make([]types.Option, len(poll.Options))
	for i, opt := range poll.Options {
		options[i] = DomainOption2Storage(opt)
	}

	tags := make([]types.Tag, len(poll.Tags))
	for i, tag := range poll.Tags {
		tags[i] = DomainTag2Storage(tag)
	}

	return types.Poll{
		Title:   poll.Title,
		Tags:    tags,
		Options: options,
	}
}
