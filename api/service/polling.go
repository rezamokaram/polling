package service

import (
	"context"
	"polling/api/pb"
	"polling/internal/polling"
	"polling/internal/polling/domain"
	pollPort "polling/internal/polling/port"
)

var (
	ErrPollNotFound = polling.ErrPollNotFound
)

type PollingService struct {
	svc pollPort.Service
}

func NewPollingService(svc pollPort.Service) *PollingService {
	return &PollingService{
		svc: svc,
	}
}

func (s *PollingService) CreatePoll(ctx context.Context, req *pb.CreatePollRequest) error {
	options := make([]domain.Option, len(req.GetOptions()))
	for i, opt := range req.GetOptions() {
		options[i] = domain.Option{
			Title: opt,
		}
	}

	tags := make([]domain.Tag, len(req.GetTags()))
	for i, tag := range req.GetTags() {
		tags[i] = domain.Tag{
			Title: tag,
		}
	}

	return s.svc.CreatePoll(ctx, domain.Poll{
		Title:   req.GetTitle(),
		Options: options,
		Tags:    tags,
	})
}

func (s *PollingService) PollList(ctx context.Context, req *pb.PollListRequest) (*pb.PollListResponse, error) {
	list, err := s.svc.PollList(ctx, domain.Filter{
		UserId: uint(req.GetUserId()),
		Page:   uint(req.GetPage()),
		Limit:  uint(req.GetLimit()),
		Tag: domain.Tag{
			Title: req.GetTag(),
		},
	})

	if err != nil {
		return &pb.PollListResponse{}, err
	}

	resp := pb.PollListResponse{
		Polls: make([]*pb.Poll, len(list)),
	}
	for i, poll := range list {
		resp.Polls[i] = &pb.Poll{
			Title:   poll.Title,
			Options: make([]string, len(poll.Options)),
			Tags:    make([]string, len(poll.Tags)),
		}
		for j, opt := range poll.Options {
			resp.Polls[i].Options[j] = opt.Title
		}
		for j, tag := range poll.Tags {
			resp.Polls[i].Tags[j] = tag.Title
		}
	}
	return &resp, nil
}
