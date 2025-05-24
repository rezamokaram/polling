package types

import "gorm.io/gorm"

type Poll struct {
	gorm.Model
	Title   string   `gorm:"not null"`
	Tags    []Tag    `gorm:"foreignKey:PollID;constraint:OnDelete:CASCADE;" json:"tags"`
	Options []Option `gorm:"foreignKey:PollID;constraint:OnDelete:CASCADE;" json:"options"`
}

func (Poll) TableName() string {
	return "polls"
}
