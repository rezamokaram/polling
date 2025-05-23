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
	svc                   pollPort.Service
	authSecret            string
	expMin, refreshExpMin uint
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
