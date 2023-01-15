package common

import "github.com/evgen1067/hw12_13_14_15_calendar/internal/repository"

type Exception struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseEventID struct {
	Code    int                `json:"code"`
	EventID repository.EventID `json:"eventId"`
}

type ResponseEventList struct {
	Code   int                `json:"code"`
	Events []repository.Event `json:"events"`
}
