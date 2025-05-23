package domain

type (
	TagId uint
)

type Tag struct {
	ID     TagId
	PollID PollID
	Title  string
}
