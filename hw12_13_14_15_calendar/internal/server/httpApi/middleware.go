package httpApi

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/evgen1067/hw12_13_14_15_calendar/internal/logger"
)

type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{w, http.StatusOK}
}

func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func headersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lrw := NewLoggingResponseWriter(w)
		next.ServeHTTP(lrw, r)
		latency := time.Since(start).Nanoseconds()
		// get client ip address
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		message := fmt.Sprintf("%v %v %v %v %v %v ns %v",
			ip,
			r.Method,
			r.URL.Path,
			r.Proto,
			lrw.statusCode,
			latency,
			r.UserAgent())

		switch {
		case lrw.statusCode < 300:
			logger.Logger.Info(message)
		case lrw.statusCode >= 300 && lrw.statusCode < 400:
			logger.Logger.Warn(message)
		default:
			logger.Logger.Error(message)
		}
	})
}
