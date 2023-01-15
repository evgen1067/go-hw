package transformer

import (
	"time"

	"github.com/evgen1067/hw12_13_14_15_calendar/api"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/repository"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func TransformEventToPb(e repository.Event) *api.Event {
	return &api.Event{
		Id:          uint64(e.ID),
		Title:       e.Title,
		Description: e.Description,
		DateStart:   &timestamp.Timestamp{Seconds: e.DateStart.Unix(), Nanos: int32(e.DateStart.Nanosecond())},
		DateEnd:     &timestamp.Timestamp{Seconds: e.DateEnd.Unix(), Nanos: int32(e.DateEnd.Nanosecond())},
		NotifyIn:    uint64(e.NotifyIn),
		OwnerId:     uint64(e.OwnerID),
	}
}

func TransformPbToEvent(e *api.Event) repository.Event {
	return repository.Event{
		ID:          repository.EventID(e.Id),
		Title:       e.Title,
		Description: e.Description,
		DateStart:   time.Unix(e.DateStart.Seconds, int64(e.DateStart.Nanos)),
		DateEnd:     time.Unix(e.DateEnd.Seconds, int64(e.DateEnd.Nanos)),
		NotifyIn:    int64(e.NotifyIn),
		OwnerID:     int64(e.OwnerId),
	}
}
