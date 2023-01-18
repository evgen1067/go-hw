package service

import (
	"context"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/common"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/repository"
	"time"
)

type Events interface {
	Create(event common.Event) (common.EventID, error)
	Update(id common.EventID, event common.Event) (common.EventID, error)
	Delete(id common.EventID) (common.EventID, error)
	DayList(startDate time.Time) ([]common.Event, error)
	WeekList(startDate time.Time) ([]common.Event, error)
	MonthList(startDate time.Time) ([]common.Event, error)
}

type Services struct {
	Events
}

func NewServices(ctx context.Context, db repository.EventsRepo) *Services {
	events := NewEventsService(ctx, db)

	return &Services{
		Events: events,
	}
}
