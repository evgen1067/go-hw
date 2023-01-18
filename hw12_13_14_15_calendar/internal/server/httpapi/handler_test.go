package httpapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/common"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/service"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/evgen1067/hw12_13_14_15_calendar/internal/config"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/repository/memory"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

var ctx context.Context

func CreateEventInRepo() (common.EventID, error) {
	ctx = context.Background()
	repo := memory.NewRepo()
	cfg, _ := config.InitConfig("../../../configs/config-test.json")
	s := service.NewServices(ctx, repo)
	InitHTTP(s, cfg)
	event := common.Event{
		Title:       "Title",
		Description: "Desc",
		DateStart:   time.Now(),
		DateEnd:     time.Now().AddDate(0, 0, 1),
		NotifyIn:    1,
		OwnerID:     1,
	}
	id, err := repo.Create(ctx, event)
	return id, err
}

func TestHelloWorld(t *testing.T) {
	t.Run("test hello world handler", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		HelloWorld(w, req)
		res := w.Result()
		defer res.Body.Close()
		data, err := io.ReadAll(res.Body)
		require.NoError(t, err)
		var ex common.Exception
		err = ex.UnmarshalJSON(data)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, ex.Code)
		require.Equal(t, "Hello, World!", ex.Message)
	})
}

func TestCustomNotFoundHandler(t *testing.T) {
	t.Run("custom not found handler", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()
		CustomNotFoundHandler(w, req)
		res := w.Result()
		defer res.Body.Close()
		data, err := io.ReadAll(res.Body)
		require.NoError(t, err)
		var ex common.Exception
		err = ex.UnmarshalJSON(data)
		require.NoError(t, err)
		require.Equal(t, http.StatusNotFound, ex.Code)
		require.Equal(t, "The page you are looking for does not exist.", ex.Message)
	})
}

func TestCreateEvent(t *testing.T) {
	t.Run("successful create event", func(t *testing.T) {
		repo := memory.NewRepo()
		ctx := context.Background()
		cfg, err := config.InitConfig("../../../configs/config-test.json")
		require.NoError(t, err)
		s := service.NewServices(ctx, repo)
		InitHTTP(s, cfg)

		eventNew := common.Event{
			Title:       "Title",
			Description: "Desc",
			DateStart:   time.Now(),
			DateEnd:     time.Now().AddDate(0, 0, 1),
			NotifyIn:    1,
			OwnerID:     1,
		}
		b := new(bytes.Buffer)
		err = json.NewEncoder(b).Encode(&eventNew)
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/events/new", CreateEvent).Methods(http.MethodPost)
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/events/new", b)
		require.NoError(t, err)
		router.ServeHTTP(recorder, req)
		var resp common.ResponseEventID
		err = json.NewDecoder(recorder.Body).Decode(&resp)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.Code)
		require.Equal(t, common.EventID(0), resp.EventID)

		weekList, err := repo.WeekList(ctx, time.Now().AddDate(0, 0, -1))
		require.NoError(t, err)
		require.Equal(t, 1, len(weekList))
	})
	t.Run("not successful create event", func(t *testing.T) {
		repo := memory.NewRepo()
		ctx := context.Background()
		cfg, err := config.InitConfig("../../../configs/config-test.json")
		require.NoError(t, err)
		s := service.NewServices(ctx, repo)
		InitHTTP(s, cfg)
		dateStart := time.Now()
		event := common.Event{
			Title:       "Title",
			Description: "Desc",
			DateStart:   dateStart,
			DateEnd:     time.Now().AddDate(0, 0, 1),
			NotifyIn:    1,
			OwnerID:     1,
		}
		id, err := repo.Create(ctx, event)
		require.NoError(t, err)

		eventNew := common.Event{
			Title:       "Title",
			Description: "Desc",
			DateStart:   dateStart,
			DateEnd:     time.Now().AddDate(0, 0, 1),
			NotifyIn:    1,
			OwnerID:     1,
		}
		b := new(bytes.Buffer)
		err = json.NewEncoder(b).Encode(&eventNew)
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/events/new", CreateEvent).Methods(http.MethodPost)
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/events/new", b)
		require.NoError(t, err)
		router.ServeHTTP(recorder, req)
		var resp common.ResponseEventID
		err = json.NewDecoder(recorder.Body).Decode(&resp)
		require.NoError(t, err)
		require.Equal(t, http.StatusBadRequest, resp.Code)
		require.Equal(t, id, resp.EventID)

		weekList, err := repo.WeekList(ctx, time.Now().AddDate(0, 0, -1))
		require.NoError(t, err)
		require.Equal(t, 1, len(weekList))
	})
}

func TestUpdateEvent(t *testing.T) {
	t.Run("successful 2-update.feature event", func(t *testing.T) {
		repo := memory.NewRepo()
		ctx := context.Background()
		cfg, err := config.InitConfig("../../../configs/config-test.json")
		require.NoError(t, err)
		s := service.NewServices(ctx, repo)
		InitHTTP(s, cfg)
		event := common.Event{
			Title:       "Title",
			Description: "Desc",
			DateStart:   time.Now(),
			DateEnd:     time.Now().AddDate(0, 0, 1),
			NotifyIn:    1,
			OwnerID:     1,
		}
		id, err := repo.Create(ctx, event)
		require.NoError(t, err)

		eventNew := common.Event{
			Title:       "New",
			Description: "New",
			DateStart:   time.Now(),
			DateEnd:     time.Now().AddDate(0, 0, 1),
			NotifyIn:    1,
			OwnerID:     1,
		}
		b := new(bytes.Buffer)
		err = json.NewEncoder(b).Encode(&eventNew)
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/events/{id}", UpdateEvent).Methods(http.MethodPut)
		req, err := http.NewRequestWithContext(ctx, http.MethodPut, fmt.Sprintf("/events/%d", id), b)
		require.NoError(t, err)
		router.ServeHTTP(recorder, req)
		var resp common.ResponseEventID
		err = json.NewDecoder(recorder.Body).Decode(&resp)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.Code)
		require.Equal(t, id, resp.EventID)

		weekList, err := repo.WeekList(ctx, time.Now().AddDate(0, 0, -1))
		require.NoError(t, err)
		require.Equal(t, 1, len(weekList))

		require.Equal(t, eventNew.Title, weekList[0].Title)
		require.Equal(t, eventNew.Description, weekList[0].Description)
	})
	t.Run("not successful 2-update.feature event", func(t *testing.T) {
		repo := memory.NewRepo()
		ctx := context.Background()
		cfg, err := config.InitConfig("../../../configs/config-test.json")
		require.NoError(t, err)
		s := service.NewServices(ctx, repo)
		InitHTTP(s, cfg)
		id := common.EventID(0)
		eventNew := common.Event{
			Title:       "Title",
			Description: "Desc",
			DateStart:   time.Now(),
			DateEnd:     time.Now().AddDate(0, 0, 1),
			NotifyIn:    1,
			OwnerID:     1,
		}
		b := new(bytes.Buffer)
		err = json.NewEncoder(b).Encode(&eventNew)
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/events/{id}", UpdateEvent).Methods(http.MethodPut)
		req, err := http.NewRequestWithContext(ctx, http.MethodPut, fmt.Sprintf("/events/%d", id), b)
		require.NoError(t, err)
		router.ServeHTTP(recorder, req)
		var resp common.Exception
		err = json.NewDecoder(recorder.Body).Decode(&resp)
		require.NoError(t, err)
		require.Equal(t, http.StatusBadRequest, resp.Code)
	})
}

func TestDeleteEvent(t *testing.T) {
	t.Run("successful delete event", func(t *testing.T) {
		id, err := CreateEventInRepo()
		require.NoError(t, err)
		recorder := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/events/{id}", DeleteEvent).Methods(http.MethodDelete)
		req, err := http.NewRequestWithContext(ctx, http.MethodDelete, fmt.Sprintf("/events/%d", id), nil)
		require.NoError(t, err)
		router.ServeHTTP(recorder, req)
		var resp common.ResponseEventID
		err = json.NewDecoder(recorder.Body).Decode(&resp)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.Code)
		require.Equal(t, id, resp.EventID)
	})
	t.Run("not successful delete event", func(t *testing.T) {
		id := common.EventID(0)
		recorder := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/events/{id}", DeleteEvent).Methods(http.MethodDelete)
		req, err := http.NewRequestWithContext(ctx, http.MethodDelete, fmt.Sprintf("/events/%d", id), nil)
		require.NoError(t, err)
		router.ServeHTTP(recorder, req)
		var resp common.Exception
		err = json.NewDecoder(recorder.Body).Decode(&resp)
		require.NoError(t, err)
		require.Equal(t, http.StatusBadRequest, resp.Code)
	})
}

func TestEventList(t *testing.T) {
	t.Run("not supported period", func(t *testing.T) {
		_, err := CreateEventInRepo()
		require.NoError(t, err)
		recorder := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/events/{period}", EventList).Methods(http.MethodGet)
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("/events/%v", "fail"), nil)
		require.NoError(t, err)
		router.ServeHTTP(recorder, req)
		var resp common.Exception
		err = json.NewDecoder(recorder.Body).Decode(&resp)
		require.NoError(t, err)
		require.Equal(t, http.StatusBadRequest, resp.Code)
		require.Equal(t, "The period specified in the request is not supported by the service.", resp.Message)
	})
	t.Run("no start date specified", func(t *testing.T) {
		_, err := CreateEventInRepo()
		require.NoError(t, err)
		recorder := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/events/{period}", EventList).Methods(http.MethodGet)
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("/events/%v", "day"), nil)
		require.NoError(t, err)
		router.ServeHTTP(recorder, req)
		var resp common.Exception
		err = json.NewDecoder(recorder.Body).Decode(&resp)
		require.NoError(t, err)
		require.Equal(t, http.StatusBadRequest, resp.Code)
		require.Equal(t, "No start date specified.", resp.Message)
	})
	t.Run("date parsing error", func(t *testing.T) {
		_, err := CreateEventInRepo()
		require.NoError(t, err)
		recorder := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/events/{period}", EventList).Methods(http.MethodGet)
		req, err := http.NewRequestWithContext(ctx,
			http.MethodGet,
			fmt.Sprintf("/events/%v?%v", "day", "date=2001-12-18"),
			nil)
		require.NoError(t, err)
		router.ServeHTTP(recorder, req)
		var resp common.Exception
		err = json.NewDecoder(recorder.Body).Decode(&resp)
		require.NoError(t, err)
		require.Equal(t, http.StatusBadRequest, resp.Code)
	})
	t.Run("success get list", func(t *testing.T) {
		_, err := CreateEventInRepo()
		require.NoError(t, err)
		recorder := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/events/{period}", EventList).Methods(http.MethodGet)
		req, err := http.NewRequestWithContext(ctx,
			http.MethodGet,
			fmt.Sprintf("/events/%v?date=%v", "week",
				time.Now().Add(-5*time.Second).Format("2006-01-02T15:04:05")), nil)
		require.NoError(t, err)
		router.ServeHTTP(recorder, req)
		var resp common.ResponseEventList
		err = json.NewDecoder(recorder.Body).Decode(&resp)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.Code)
		require.Equal(t, common.EventID(0), resp.Events[0].ID)
	})
}
