package httpApi

import (
	"encoding/json"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/repository"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/server/httpApi/common"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func CustomNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	ex := common.Exception{
		Code:    http.StatusNotFound,
		Message: "The page you are looking for does not exist.",
	}
	WriteException(w, ex)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	ex := common.Exception{
		Code:    http.StatusOK,
		Message: "Hello, World!",
	}
	WriteException(w, ex)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event repository.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		WriteException(w, common.Exception{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	var eventID repository.EventID
	eventID, err = restApi.repo.Create(restApi.ctx, event)
	if err != nil {
		WriteException(w, common.Exception{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	WriteEventIDResponse(w, common.ResponseEventID{
		Code:    http.StatusOK,
		EventID: eventID,
	})
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	eventID := repository.EventID(id)
	if err != nil {
		WriteException(w, common.Exception{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	var event repository.Event
	err = json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		WriteException(w, common.Exception{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	eventID, err = restApi.repo.Update(restApi.ctx, eventID, event)
	if err != nil {
		WriteException(w, common.Exception{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	WriteEventIDResponse(w, common.ResponseEventID{
		Code:    http.StatusOK,
		EventID: eventID,
	})
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		WriteException(w, common.Exception{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	eventID := repository.EventID(id)
	eventID, err = restApi.repo.Delete(restApi.ctx, eventID)
	if err != nil {
		WriteException(w, common.Exception{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	WriteEventIDResponse(w, common.ResponseEventID{
		Code:    http.StatusOK,
		EventID: eventID,
	})
}

func EventList(w http.ResponseWriter, r *http.Request) {
	_period := strings.ToLower(mux.Vars(r)["period"])
	var period repository.Period
	switch _period {
	case "day":
		period = repository.Period("Day")
	case "week":
		period = repository.Period("Week")
	case "month":
		period = repository.Period("Month")
	default:
		WriteException(w, common.Exception{
			Code:    http.StatusBadRequest,
			Message: "The period specified in the request is not supported by the service.",
		})
		return
	}
	dateParam := r.URL.Query().Get("date")
	if dateParam == "" {
		WriteException(w, common.Exception{
			Code:    http.StatusBadRequest,
			Message: "No start date specified.",
		})
		return
	}
	startDate, err := time.Parse("2006-01-02T15:04:05", dateParam)
	if err != nil {
		WriteException(w, common.Exception{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	var events []repository.Event
	switch period {
	case "Day":
		events, err = restApi.repo.DayList(restApi.ctx, startDate)
	case "Week":
		events, err = restApi.repo.WeekList(restApi.ctx, startDate)
	case "Month":
		events, err = restApi.repo.MonthList(restApi.ctx, startDate)
	}
	if err != nil {
		WriteException(w, common.Exception{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	WriteEventListResponse(w, common.ResponseEventList{
		Code:   http.StatusOK,
		Events: events,
	})
}

func WriteEventIDResponse(w http.ResponseWriter, r common.ResponseEventID) {
	w.WriteHeader(r.Code)
	jsonResponse, err := r.MarshalJSON()
	if err != nil {
		return
	}
	_, err = w.Write(jsonResponse)
	if err != nil {
		return
	}
}

func WriteEventListResponse(w http.ResponseWriter, r common.ResponseEventList) {
	w.WriteHeader(r.Code)
	jsonResponse, err := r.MarshalJSON()
	if err != nil {
		return
	}
	_, err = w.Write(jsonResponse)
	if err != nil {
		return
	}
}

func WriteException(w http.ResponseWriter, ex common.Exception) {
	w.WriteHeader(ex.Code)
	jsonResponse, err := ex.MarshalJSON()
	if err != nil {
		return
	}
	_, err = w.Write(jsonResponse)
	if err != nil {
		return
	}
}
