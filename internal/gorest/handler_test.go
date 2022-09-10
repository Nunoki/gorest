package gorest

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

const testUserID = "f44fe12d-8bec-4720-845e-dbebcc053f9e"

func TestValidPutRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := &RepositoryMock{}
	s := Server{
		router: gin.New(),
	}

	handler := NewHandler(repo)
	s.router.PUT("/", func(c *gin.Context) {
		c.Set("userID", testUserID)
		handler.HandleStore(c)
	})

	repo.UpdateFunc = func(uid string, content []byte) error {
		return nil
	}
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PUT", "/", strings.NewReader(`"hello world"`))
	r.Header.Set("Content-type", "application/json")
	s.ServeHTTP(w, r)

	res := w.Result()

	if res.StatusCode != 200 {
		t.Fatalf("PUT request with valid JSON has failed with code %d", res.StatusCode)
	}
}

func TestPutRequestWithRepositoryErrorShouldFailWith500(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := &RepositoryMock{}
	s := Server{
		router: gin.New(),
	}

	handler := NewHandler(repo)
	s.router.PUT("/", func(c *gin.Context) {
		c.Set("userID", testUserID)
		handler.HandleStore(c)
	})

	repo.UpdateFunc = func(uid string, content []byte) error {
		return errors.New("error")
	}
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PUT", "/", strings.NewReader(`{"message":"Hello"}`))
	r.Header.Set("Content-type", "application/json")
	s.ServeHTTP(w, r)

	res := w.Result()

	if res.StatusCode != 500 {
		t.Fatal("Repository error should result in 500 response, instead received", res.StatusCode)
	}
}

func TestSavingEmptyPayloadReturns400(t *testing.T) {
	gin.SetMode(gin.TestMode)
	s := Server{
		router: gin.New(),
	}

	repo := &RepositoryMock{}
	handler := NewHandler(repo)
	s.router.PUT("/", func(c *gin.Context) {
		c.Set("userID", testUserID)
		handler.HandleStore(c)
	})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PUT", "/", strings.NewReader(""))
	r.Header.Set("Content-type", "application/json")
	s.ServeHTTP(w, r)

	res := w.Result()

	if res.StatusCode != 400 {
		t.Fatalf("Empty payload should result in error 400, received %d instead", res.StatusCode)
	}
}

func TestReturnedDateFormat(t *testing.T) {
	time_, err := time.Parse("January 2, 15:04:05, 2006", "January 2, 15:04:05, 2006")
	if err != nil {
		t.Fatal("Failed to parse date")
	}

	gin.SetMode(gin.TestMode)
	s := Server{
		router: gin.New(),
	}

	repo := &RepositoryMock{}
	repo.FindFunc = func(string) ([]byte, time.Time, error) {
		return []byte(""), time_, nil
	}
	handler := NewHandler(repo)
	s.router.GET("/", func(c *gin.Context) {
		c.Set("userID", testUserID)
		handler.HandleRead(c)
	})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", strings.NewReader(""))
	r.Header.Set("Content-type", "application/json")
	s.ServeHTTP(w, r)

	exp := "2006-01-02T15:04:05Z"
	if !strings.Contains(w.Body.String(), exp) {
		t.Fatalf("Date format expected to be %q in response %q", exp, w.Body.String())
	}
}
