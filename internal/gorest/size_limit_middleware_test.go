package gorest

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestValidSizePayload(t *testing.T) {
	// log.SetOutput(ioutil.Discard)
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(SizeLimitMiddleware(6))

	handlerOK := false
	router.PUT("/", func(c *gin.Context) {
		handlerOK = true
	})

	resp := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/", strings.NewReader("12345"))
	req.Header.Set("Content-type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("Expected response code %d, got %d", http.StatusOK, resp.Code)
	}

	if !handlerOK {
		t.Fatal("Handler didn't get called")
	}
}

func TestFailsWithTooLargePayload(t *testing.T) {
	// log.SetOutput(ioutil.Discard)
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(SizeLimitMiddleware(5))

	router.PUT("/", func(c *gin.Context) {
		t.Fatal("handler shouldn't be called when payload too large")
	})

	resp := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/", strings.NewReader("123456"))
	req.Header.Set("Content-type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusRequestEntityTooLarge {
		t.Fatalf("Expected response code %d, got %d", http.StatusRequestEntityTooLarge, resp.Code)
	}
}
