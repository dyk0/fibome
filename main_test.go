package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

//Test imputing negative number returns error

func TestNegativeNumber(t *testing.T) {
	t.Log("When the number given is negative")
	request, err := http.NewRequest("GET", "/-10", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := newRequestRecorder(request, "GET", "/:size", fib)
	if recorder.Code != 400 {
		t.Error("Expected response code to be 400")
	}

	expected_response := "Invalid Entry! Please enter a positive integer\n"
	if recorder.Body.String() != expected_response {
		t.Error("Response does not match")
	}
}

func TestNaN(t *testing.T) {
	t.Log("When the number is NaN")
	request, err := http.NewRequest("GET", "/fibonocci", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := newRequestRecorder(request, "GET", "/:size", fib)
	if recorder.Code != 400 {
		t.Error("Expected response code to be 400")
	}

	expected_response := "Invalid Entry! Please enter a positive integer\n"
	if recorder.Body.String() != expected_response {
		t.Error("Response does not match")
	}
}
func TestPositiveNumber(t *testing.T) {
	t.Log("When the number is positive")
	request, err := http.NewRequest("GET", "/5", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := newRequestRecorder(request, "GET", "/:size", fib)
	if recorder.Code != 200 {
		t.Error("Expected response code to be 200")
	}

	expected_response := "[0 1 1 2 3]\n"
	if recorder.Body.String() != expected_response {
		t.Error("Response does not match")
	}
}

// Mocks a handler and returns a httptest.ResponseRecorder
func newRequestRecorder(req *http.Request, method string, strPath string, fnHandler func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) *httptest.ResponseRecorder {
	router := httprouter.New()
	router.Handle(method, strPath, fnHandler)
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	router.ServeHTTP(rr, req)
	return rr
}
