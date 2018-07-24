package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func getRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// Test the main file
func TestMain(t *testing.T) {
	go main()
}

// Test that a GET request to the home page returns the status page
func TestIndexPage(t *testing.T) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	r := getRouter()

	// Define the route similar to its definition in the routes file
	r.GET("/", Index)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	// Test that the http status code is 200
	if w.Code != http.StatusOK {
		t.Fail()
	}

	// Test that the page contains `time`
	p, err := ioutil.ReadAll(w.Body)
	if err != nil || strings.Index(string(p), "uptime") < 0 {
		t.Fail()
	}
}


// Test that a GET returns a 404 when an invalid URL is requested
func Test404Page(t *testing.T) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	r := getRouter()

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/qwerty", nil)

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	// Test that the http status code is 404
	if w.Code != http.StatusNotFound {
		t.Fail()
	}

	// Test that the page contains `404`
	p, err := ioutil.ReadAll(w.Body)
	if err != nil || strings.Index(string(p), "404") < 0 {
		t.Fail()
	}
}

// Test that a favicon.ico request returns a 404
func TestFaviconNotFound(t *testing.T) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	r := getRouter()

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/favicon.ico", nil)

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	// Test that the http status code is 404
	if w.Code != http.StatusNotFound {
		t.Fail()
	}

	// Test that the page contains `404`
	p, err := ioutil.ReadAll(w.Body)
	if err != nil || strings.Index(string(p), "404") < 0 {
		t.Fail()
	}
}
