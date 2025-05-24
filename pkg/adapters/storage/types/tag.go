package types

import "gorm.io/gorm"

type Tag struct {
	gorm.Model

	Title  string `gorm:"not null"`
	PollID uint   `gorm:"not null"`
}

func (Tag) TableName() string {
	return "tags"
}
