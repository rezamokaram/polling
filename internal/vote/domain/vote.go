package domain

import "errors"

var (
	ErrAlreadyVoted   = errors.New("already voted")
	ErrOptionNotFound = errors.New("option not found")
	ErrPollNotFound   = errors.New("poll not found")
	ErrVoteNotFound   = errors.New("vote not found")
)

type Vote struct {
	PollID uint `json:"poll_id"`
	UserID uint `json:"user_id"`
	Index  uint `json:"index"`
}
