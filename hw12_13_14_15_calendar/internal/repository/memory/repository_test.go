package memory

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/evgen1067/hw12_13_14_15_calendar/internal/repository"
	"github.com/stretchr/testify/require"
)

func TestMemoryRepo(t *testing.T) {
	t.Run("test memory repository", func(t *testing.T) {
		var err error
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		repo := NewRepo()
		events := make([]repository.Event, 0)
		for i := 0; i < 10; i++ {
			e := repository.Event{
				Title:       fmt.Sprintf("Title %v", i),
				Description: fmt.Sprintf("Description %v", i),
				DateStart:   time.Now().AddDate(0, 0, i*1),
				DateEnd:     time.Now().AddDate(0, 0, i*2),
			}
			events = append(events, e)
		}
		for _, e := range events {
			_, err = repo.Create(ctx, e)
			require.NoError(t, err)
		}

		_, err = repo.Create(ctx, events[2])
		require.Error(t, err)
		require.ErrorIs(t, ErrDateBusy, err)

		_, err = repo.Update(ctx, 3, events[2])
		require.Error(t, err)
		require.ErrorIs(t, ErrDateBusy, err)

		_, err = repo.Delete(ctx, 0)
		require.NoError(t, err)

		_, err = repo.Delete(ctx, 0)
		require.Error(t, err)

		e := repository.Event{
			Title:       fmt.Sprintf("Title %v", 2),
			Description: fmt.Sprintf("Description %v", 2),
			DateStart:   time.Now().AddDate(0, 0, 1),
			DateEnd:     time.Now().AddDate(0, 0, 2),
		}

		_, err = repo.Update(ctx, 0, e)
		require.Error(t, err)

		_, err = repo.Update(ctx, 1, e)
		require.NoError(t, err)

		var periodEvents []repository.Event

		periodEvents, err = repo.DayList(ctx, time.Now().AddDate(0, 0, 1))
		require.NoError(t, err)
		require.Equal(t, 2, len(periodEvents))

		periodEvents, err = repo.DayList(ctx, time.Now().AddDate(0, 0, -30))
		require.NoError(t, err)
		require.Equal(t, 0, len(periodEvents))

		periodEvents, err = repo.WeekList(ctx, time.Now().AddDate(0, 0, 1))
		require.NoError(t, err)
		require.Equal(t, 8, len(periodEvents))

		periodEvents, err = repo.MonthList(ctx, time.Now().AddDate(0, 0, 1))
		require.NoError(t, err)
		require.Equal(t, 9, len(periodEvents))
	})
}
