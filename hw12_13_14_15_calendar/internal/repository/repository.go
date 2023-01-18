package repository

import (
	"context"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/common"
	"time"
)

type Period string

type DatabaseRepo interface {
	Connect(ctx context.Context) error
	Close() error
	EventsRepo
	SchedulerRepo
}

type EventsRepo interface {
	Create(ctx context.Context, event common.Event) (common.EventID, error)
	Update(ctx context.Context, id common.EventID, event common.Event) (common.EventID, error)
	Delete(ctx context.Context, id common.EventID) (common.EventID, error)
	DayList(ctx context.Context, startDate time.Time) ([]common.Event, error)
	WeekList(ctx context.Context, startDate time.Time) ([]common.Event, error)
	MonthList(ctx context.Context, startDate time.Time) ([]common.Event, error)
}

type SchedulerRepo interface {
	SchedulerList(ctx context.Context) ([]common.Notice, error)
	ClearOldEvents(ctx context.Context) error
}
