package service

import (
	"context"
	"polling/api/pb"
	"polling/internal/vote/domain"
	"polling/internal/vote/port"
)

type VoteService struct {
	svc port.Service
}

func NewVoteService(svc port.Service) *VoteService {
	return &VoteService{
		svc: svc,
	}
}

func (s *VoteService) VotePoll(ctx context.Context, req *pb.VotePollRequest) error {
	return s.svc.VotePoll(ctx, domain.Vote{
		UserID: uint(req.GetUserId()),
		PollID: uint(req.GetPollId()),
		Index:  uint(req.GetOptionIndex()),
	})
}
