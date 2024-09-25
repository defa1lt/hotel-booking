package middleware

import (
	"github.com/gorilla/mux"
	"net/http"
	"runtime/debug"

	"github.com/sirupsen/logrus"
)

func RecoverMiddleware(log *logrus.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					log.Errorf("Паника: %v\n%s", err, debug.Stack())
					http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
