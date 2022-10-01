package presentation

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		begin := time.Now()

		next.ServeHTTP(w, r)

		exectime := time.Since(begin)

		logrus.WithFields(logrus.Fields{
			"method":   r.Method,
			"url":      r.URL,
			"duration": exectime.Milliseconds(),
		}).WithTime(time.Now()).Info()
	})
}
