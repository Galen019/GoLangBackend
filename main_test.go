package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Start the server in a goroutine
	go startServer()

	// Give the server a moment to start
	time.Sleep(1 * time.Second)

	// Run the tests
	m.Run()
}

func TestDeleteSubscription(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Set up the router
	router := setupRouter()

	// Create a new HTTP request for deleting a subscription
	req, err := http.NewRequest("DELETE", "/subscriptions/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP response recorder
	recorder := httptest.NewRecorder()

	// Serve the HTTP request
	router.ServeHTTP(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status OK; got %d", recorder.Code)
	}

	// Check the response body
	expectedBody := `{"id":1,"user_id":0,"plan":"","active":false,"amount":0}`
	if recorder.Body.String() != expectedBody {
		t.Errorf("expected body %s; got %s", expectedBody, recorder.Body.String())
	}
}

func TestGetSubscriptions(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Set up the router
	router := setupRouter()

	// Create a new HTTP request for deleting a subscription
	req, err := http.NewRequest("GET", "/subscriptions", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP response recorder
	recorder := httptest.NewRecorder()

	// Serve the HTTP request
	router.ServeHTTP(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status OK; got %d", recorder.Code)
	}

	// Check the response body
	fmt.Println(recorder.Body.String())
	if recorder.Body.String() == "" {
		t.Errorf("Response is empty %s", recorder.Body.String())
	}
}
