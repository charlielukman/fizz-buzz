package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFizzBuzzHandler(t *testing.T) {
	testFizzBuzzHandler_paramRange_more_than_100(t)
	testFizzBuzzHandler_param_from_greater_than_to(t)
	testFizzBuzzHandler_param_from_not_number(t)
	testFizzBuzzHandler_param_to_not_number(t)
	testFizzBuzzHandler(t)
}

func testFizzBuzzHandler_paramRange_more_than_100(t *testing.T) {
	req, err := http.NewRequest("GET", "/range-fizzbuzz?from=1&to=200", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(fizzBuzzHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expected := "Invalid parameter from and to more than 100"
	if strings.TrimSuffix(rr.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			strings.TrimSuffix(rr.Body.String(), "\n"), expected)
	}
}

func testFizzBuzzHandler_param_from_greater_than_to(t *testing.T) {
	req, err := http.NewRequest("GET", "/range-fizzbuzz?from=10&to=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(fizzBuzzHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expected := "Invalid parameter value from > to"
	if strings.TrimSuffix(rr.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			strings.TrimSuffix(rr.Body.String(), "\n"), expected)
	}
}

func testFizzBuzzHandler_param_from_not_number(t *testing.T) {
	req, err := http.NewRequest("GET", "/range-fizzbuzz?from=a&to=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(fizzBuzzHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expected := "Invalid parameter value from"
	if strings.TrimSuffix(rr.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			strings.TrimSuffix(rr.Body.String(), "\n"), expected)
	}
}

func testFizzBuzzHandler_param_to_not_number(t *testing.T) {
	req, err := http.NewRequest("GET", "/range-fizzbuzz?from=1&to=b", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(fizzBuzzHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expected := "Invalid parameter value to"
	if strings.TrimSuffix(rr.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			strings.TrimSuffix(rr.Body.String(), "\n"), expected)
	}
}

func testFizzBuzzHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/range-fizzbuzz?from=1&to=15", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(fizzBuzzHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz"
	if strings.TrimSuffix(rr.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			strings.TrimSuffix(rr.Body.String(), "\n"), expected)
	}
}
