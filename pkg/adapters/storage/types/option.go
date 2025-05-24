package types

import "gorm.io/gorm"

type Option struct {
	gorm.Model

	Title  string `gorm:"not null"`
	PollID uint   `gorm:"not null"`
	Index  uint   `gorm:"not null"`
}

func (Option) TableName() string {
	return "options"
}
