package mapper

type PollStats struct {
	PollID     uint
	VotesStats []VoteStat
}

type VoteStat struct {
	OptionTitle string
	VotesCount  uint
}
