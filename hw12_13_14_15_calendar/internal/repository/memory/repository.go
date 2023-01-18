package memory

import (
	"context"
	"errors"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/common"
	"sync"
	"time"

	"github.com/evgen1067/hw12_13_14_15_calendar/internal/repository"
)

var (
	ErrNotFound = errors.New("event not found")
	ErrDateBusy = errors.New("this time is already occupied by another event")
)

type Repo struct {
	mu        sync.RWMutex
	Events    map[common.EventID]common.Event
	Increment common.EventID
	length    int
}

func NewRepo() repository.EventsRepo {
	return &Repo{
		Events: make(map[common.EventID]common.Event),
	}
}

func (r *Repo) CheckDate(ctx context.Context, event common.Event) error {
	for _, e := range r.Events {
		if e.DateStart.Format("2/Jan/2006:15:04") == event.DateStart.Format("2/Jan/2006:15:04") && e.ID != event.ID {
			return ErrDateBusy
		}
	}
	return nil
}

func (r *Repo) Create(ctx context.Context, event common.Event) (common.EventID, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	event.ID = r.Increment
	err := r.CheckDate(ctx, event)
	if err != nil {
		return event.ID, err
	}
	r.Events[event.ID] = event
	r.Increment++
	r.length++
	return event.ID, nil
}

func (r *Repo) Update(ctx context.Context, id common.EventID, event common.Event) (common.EventID, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.Events[id]
	if !ok {
		return id, ErrNotFound
	}

	event.ID = id

	err := r.CheckDate(ctx, event)
	if err != nil {
		return event.ID, err
	}
	r.Events[id] = event
	return event.ID, nil
}

func (r *Repo) Delete(ctx context.Context, id common.EventID) (common.EventID, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.Events[id]
	if !ok {
		return id, ErrNotFound
	}
	delete(r.Events, id)
	r.length--
	return id, nil
}

func (r *Repo) PeriodList(
	ctx context.Context,
	startPeriod time.Time,
	period repository.Period,
) ([]common.Event, error) {
	var endPeriod time.Time
	switch period {
	case "Day":
		endPeriod = startPeriod.AddDate(0, 0, 1)
	case "Week":
		endPeriod = startPeriod.AddDate(0, 0, 7)
	case "Month":
		endPeriod = startPeriod.AddDate(0, 1, 0)
	}
	var events []common.Event
	for _, e := range r.Events {
		if e.DateEnd.After(startPeriod) && endPeriod.After(e.DateStart) {
			events = append(events, e)
		}
	}
	return events, nil
}

func (r *Repo) DayList(ctx context.Context, startDate time.Time) ([]common.Event, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	period := repository.Period("Day")
	return r.PeriodList(ctx, startDate, period)
}

func (r *Repo) WeekList(ctx context.Context, startDate time.Time) ([]common.Event, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	period := repository.Period("Week")
	return r.PeriodList(ctx, startDate, period)
}

func (r *Repo) MonthList(ctx context.Context, startDate time.Time) ([]common.Event, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	period := repository.Period("Month")
	return r.PeriodList(ctx, startDate, period)
}
