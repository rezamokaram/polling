package mapper

import (
	"polling/internal/polling/domain"
	"polling/pkg/adapters/storage/types"
)

func DomainTag2Storage(tag domain.Tag) types.Tag {
	return types.Tag{
		Title: tag.Title,
	}
}

func StorageTag2Domain(tag types.Tag) domain.Tag {
	return domain.Tag{
		ID:    domain.TagId(tag.ID),
		Title: tag.Title,
	}
}
