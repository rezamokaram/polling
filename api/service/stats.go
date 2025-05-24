package service

import (
	"context"
	"log"
	"polling/api/pb"
	"polling/internal/aggregates/stats"
	"polling/internal/polling/domain"
)

type StatsService struct {
	svc stats.AggregatePollStatsService
}

func NewStatsService(svc stats.AggregatePollStatsService) *StatsService {
	return &StatsService{
		svc: svc,
	}
}

func (s *StatsService) PollStats(ctx context.Context, req *pb.PollStatsRequest) (*pb.PollStatsResponse, error) {
	pollStats, err := s.svc.GetPollStats(ctx, domain.PollID(req.GetPollId()))
	if err != nil {
		log.Println("error getting poll stats:", err)
		return nil, err
	}
	resp := &pb.PollStatsResponse{
		PollId: uint32(pollStats.PollID),
		Votes:  make([]*pb.VoteStats, len(pollStats.VotesStats)),
	}
	for i, stat := range pollStats.VotesStats {
		resp.Votes[i] = &pb.VoteStats{
			Option: stat.OptionTitle,
			Count:  uint32(stat.VotesCount),
		}
	}

	return resp, nil
}
