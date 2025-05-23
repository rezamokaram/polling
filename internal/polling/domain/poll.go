package domain

import "time"

type (
	PollID uint
)

type Poll struct {
	ID        PollID
	CreatedAt time.Time
	Title     string
	Options   []Option
	Tags      []Tag
}
