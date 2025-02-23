package middleware

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func LoggingMiddleware(log *logrus.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			duration := time.Since(start)
			log.Infof("%s %s %s", r.Method, r.RequestURI, duration)
		})
	}
}
