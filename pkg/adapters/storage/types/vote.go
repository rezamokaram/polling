package types

import "gorm.io/gorm"

type Vote struct {
	gorm.Model

	PollID   uint `json:"poll_id"`
	UserID   uint `json:"user_id"`
	OptionID uint `json:"option_id"`
}

func (Vote) TableName() string {
	return "votes"
}
