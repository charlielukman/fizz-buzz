package middleware

import (
	"bytes"
	"log"
	"net/http"
	"time"
)

const MAX_CONCURRENT_REQUESTS = 1000

var (
	sem = make(chan struct{}, MAX_CONCURRENT_REQUESTS)
)

type responseRecorder struct {
	http.ResponseWriter
	status int
	body   *bytes.Buffer
}

func (rr *responseRecorder) WriteHeader(status int) {
	rr.status = status
	rr.ResponseWriter.WriteHeader(status)
}

func (rr *responseRecorder) Write(body []byte) (int, error) {
	rr.body.Write(body)
	return rr.ResponseWriter.Write(body)
}

func LimitConcurrentRequestsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case sem <- struct{}{}:
			defer func() { <-sem }()
		default:
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func TimeoutMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeoutHandler := http.TimeoutHandler(next, 1*time.Second, "Request timeout")
		timeoutHandler.ServeHTTP(w, r)
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rr := &responseRecorder{
			ResponseWriter: w,
			body:           &bytes.Buffer{},
		}

		startTime := time.Now()
		next.ServeHTTP(rr, r)
		duration := time.Since(startTime)
		log.Printf("Request: %s, Response Status: %d, Response Body: %s, Latency: %s\n",
			r.URL.String(), rr.status, rr.body.String(), duration)
	})
}
