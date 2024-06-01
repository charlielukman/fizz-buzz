package main

import (
	"context"
	"fizz-buzz/fizzbuzz"
	"fizz-buzz/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func fizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	const maxRange = 100

	fromStr := r.URL.Query().Get("from")
	toStr := r.URL.Query().Get("to")

	from, err := strconv.Atoi(fromStr)
	if err != nil {
		http.Error(w, "Invalid parameter value from", http.StatusBadRequest)
		return
	}

	to, err := strconv.Atoi(toStr)
	if err != nil {
		http.Error(w, "Invalid parameter value to", http.StatusBadRequest)
		return
	}

	if from > to {
		http.Error(w, "Invalid parameter value from > to", http.StatusBadRequest)
		return
	}

	if to-from > maxRange-1 {
		http.Error(w, "Invalid parameter from and to more than 100", http.StatusBadRequest)
		return
	}

	results := fizzbuzz.RangeFizzBuzz(from, to)

	response := strings.Join(results, " ")
	w.Write([]byte(response))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/range-fizzbuzz", func(w http.ResponseWriter, r *http.Request) {
		fizzBuzzHandler(w, r)
	})

	h := middleware.LimitConcurrentRequestsMiddleware(mux)
	h = middleware.TimeoutMiddleware(h)
	h = middleware.LoggingMiddleware(h)

	server := &http.Server{
		Addr:    ":8080",
		Handler: h,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on :8080: %v\n", err)
		}
	}()
	log.Println("Server is running on port 8080")

	singnalChan := make(chan os.Signal, 1)
	signal.Notify(singnalChan, os.Interrupt, syscall.SIGTERM)

	<-singnalChan

	log.Println("Shutting down the server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server Shutdown Failed:%+v", err)
	}
	log.Println("Server exited properly")
}
