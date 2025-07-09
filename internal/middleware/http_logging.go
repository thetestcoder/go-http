package middleware

import (
	"log"
	"net/http"
	"time"
)

func HttpLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrw := &responseWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}

		defer func() {
			log.Printf("%s %s %s %d %s",
				r.Method,
				r.RequestURI,
				r.Proto,
				wrw.status,
				time.Since(start),
			)
		}()

		next.ServeHTTP(wrw, r)

	})
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}
