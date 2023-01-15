package httpapi

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/evgen1067/hw12_13_14_15_calendar/internal/config"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/repository"
	"github.com/gorilla/mux"
)

var restAPI *Server

type Deps struct {
	ctx  context.Context
	repo repository.EventsRepo
}

type Server struct {
	Deps
	Srv *http.Server
}

func HTTPRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", HelloWorld).Methods(http.MethodGet)
	router.HandleFunc("/events/new", CreateEvent).Methods(http.MethodPost)
	router.HandleFunc("/events/{id}", UpdateEvent).Methods(http.MethodPut)
	router.HandleFunc("/events/{id}", DeleteEvent).Methods(http.MethodDelete)
	router.HandleFunc("/events/list/{period}", EventList).Methods(http.MethodGet)

	router.NotFoundHandler = router.NewRoute().HandlerFunc(CustomNotFoundHandler).GetHandler()

	router.Use(headersMiddleware)
	router.Use(loggerMiddleware)

	return router
}

func InitHTTP(_ctx context.Context, _repo repository.EventsRepo, cfg *config.Config) *Server {
	restAPI = &Server{
		Deps: Deps{
			ctx:  _ctx,
			repo: _repo,
		},
		Srv: &http.Server{
			Addr:              net.JoinHostPort(cfg.HTTP.Host, cfg.HTTP.Port),
			Handler:           HTTPRouter(),
			ReadHeaderTimeout: 1 * time.Second,
		},
	}
	return restAPI
}
