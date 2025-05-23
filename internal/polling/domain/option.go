package domain

type (
	OptionId uint
)

type Option struct {
	ID     OptionId
	PollID PollID
	Title  string
}
