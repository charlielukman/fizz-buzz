package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestTimeoutMiddleware(t *testing.T) {
	testTimeout(t)
	testNoTimeout(t)
}

func testTimeout(t *testing.T) {
	slowHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
	})

	handler := TimeoutMiddleware(slowHandler)

	req, err := http.NewRequest("GET", "/slow", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusServiceUnavailable {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusServiceUnavailable)
	}

	expected := "Request timeout"
	if strings.TrimSuffix(rr.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			strings.TrimSuffix(rr.Body.String(), "\n"), expected)
	}
}

func testNoTimeout(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(([]byte)("OK"))
	})

	handler := TimeoutMiddleware(h)

	req, err := http.NewRequest("GET", "/no-timeout", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestLimitConcurrentRequestsMiddleware(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Millisecond)
		w.Write([]byte("OK"))
	})

	handler := LimitConcurrentRequestsMiddleware(h)

	var wg sync.WaitGroup

	totalConcurrentRequests := MAX_CONCURRENT_REQUESTS + 1
	responses := make(chan int, totalConcurrentRequests)

	for i := 0; i < totalConcurrentRequests; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			rr := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/", nil)

			handler.ServeHTTP(rr, req)

			responses <- rr.Code
		}()
	}

	wg.Wait()

	close(responses)

	isTooManyRequests := false
	for res := range responses {
		if res == http.StatusTooManyRequests {
			isTooManyRequests = true
		}
	}

	if !isTooManyRequests {
		t.Errorf("Expected some requests to be limited, but none were")
	}
}
