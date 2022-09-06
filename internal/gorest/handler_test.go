package gorest

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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

// TODO: test that empty payload returns a 400 error
// TODO: test output format and the date format
