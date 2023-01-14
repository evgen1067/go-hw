package grpcApi

import (
	"context"
	"github.com/evgen1067/hw12_13_14_15_calendar/api"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/repository"
	data "github.com/evgen1067/hw12_13_14_15_calendar/internal/server/grpcApi/transformer"
	"time"
)

func (s *Server) Create(ctx context.Context, req *api.CreateRequest) (*api.CreateResponse, error) {
	e := data.TransformPbToEvent(req.Event)

	eId, err := s.repo.Create(ctx, e)
	if err != nil {
		return nil, err
	}

	return &api.CreateResponse{
		Id: uint64(eId),
	}, nil
}

func (s *Server) Update(ctx context.Context, req *api.UpdateRequest) (*api.UpdateResponse, error) {
	id := repository.EventID(req.Id)
	e := data.TransformPbToEvent(req.Event)

	eId, err := s.repo.Update(ctx, id, e)
	if err != nil {
		return nil, err
	}

	return &api.UpdateResponse{
		Id: uint64(eId),
	}, nil
}

func (s *Server) Delete(ctx context.Context, req *api.DeleteRequest) (*api.DeleteResponse, error) {
	id := repository.EventID(req.Id)

	eId, err := s.repo.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	return &api.DeleteResponse{
		Id: uint64(eId),
	}, nil
}

func (s *Server) DayList(ctx context.Context, req *api.ListRequest) (*api.ListResponse, error) {
	return PeriodList(ctx, req, s.repo.DayList)
}

func (s *Server) WeekList(ctx context.Context, req *api.ListRequest) (*api.ListResponse, error) {
	return PeriodList(ctx, req, s.repo.WeekList)
}

func (s *Server) MonthList(ctx context.Context, req *api.ListRequest) (*api.ListResponse, error) {
	return PeriodList(ctx, req, s.repo.MonthList)
}

func PeriodList(ctx context.Context,
	req *api.ListRequest,
	fn func(ctx context.Context, startDate time.Time) ([]repository.Event, error),
) (*api.ListResponse, error) {
	startDate := time.Unix(req.Date.Seconds, int64(req.Date.Nanos))
	events, err := fn(ctx, startDate)
	if err != nil {
		return nil, err
	}
	periodList := make([]*api.Event, 0)
	for _, val := range events {
		periodList = append(periodList, data.TransformEventToPb(val))
	}
	return &api.ListResponse{Event: periodList}, nil
}
