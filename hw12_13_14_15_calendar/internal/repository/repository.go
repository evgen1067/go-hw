package repository

import (
	"context"
	"time"
)

type EventID int64

type Event struct {
	ID          EventID   `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DateStart   time.Time `json:"dateStart"`
	DateEnd     time.Time `json:"dateEnd"`
	NotifyIn    int64     `json:"notifyIn"`
	OwnerID     int64     `json:"ownerId"`
}

type Notice struct {
	EventID  EventID   `json:"eventId"`
	Title    string    `json:"title"`
	Datetime time.Time `json:"datetime"`
	OwnerID  int64     `json:"ownerId"`
}

type Period string

type DatabaseRepo interface {
	Connect(ctx context.Context) error
	Close() error
	EventsRepo
}

type EventsRepo interface {
	Create(ctx context.Context, event Event) (EventID, error)
	Update(ctx context.Context, id EventID, event Event) (EventID, error)
	Delete(tx context.Context, id EventID) (EventID, error)
	DayList(ctx context.Context, startDate time.Time) ([]Event, error)
	WeekList(ctx context.Context, startDate time.Time) ([]Event, error)
	MonthList(ctx context.Context, startDate time.Time) ([]Event, error)
}
