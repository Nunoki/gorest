package gorest

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/nunoki/gorest/internal/gorest/middleware"
)

const testUserID = "f44fe12d-8bec-4720-845e-dbebcc053f9e"

func TestPut(t *testing.T) {
	req, err := http.NewRequest("PUT", "/", strings.NewReader("123"))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-type", "application/json")

	ctx := context.WithValue(req.Context(), middleware.UserID, testUserID)
	req = req.WithContext(ctx)

	repoMock := &RepositoryMock{}
	repoMock.UpdateFunc = func(uid string, content []byte) error {
		return nil
	}
	h := NewHandler(repoMock)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.handlePut)

	handler.ServeHTTP(rr, req)

	expCode := http.StatusOK
	if rr.Code != expCode {
		t.Errorf("expected status code %d but got %d", expCode, rr.Code)
	}

	expectedBody := "3\n"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rr.Body.String())
	}
}

func TestPutReturns500WhenRepositoryError(t *testing.T) {
	req, err := http.NewRequest("PUT", "/", strings.NewReader("123"))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-type", "application/json")

	ctx := context.WithValue(req.Context(), middleware.UserID, testUserID)
	req = req.WithContext(ctx)

	repoMock := &RepositoryMock{}
	repoMock.UpdateFunc = func(uid string, content []byte) error {
		return errors.New("i failed")
	}
	h := NewHandler(repoMock)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.handlePut)

	handler.ServeHTTP(rr, req)

	expCode := http.StatusInternalServerError
	if rr.Code != expCode {
		t.Errorf("expected status code %d but got %d", expCode, rr.Code)
	}
}

func TestPutReturns400WhenSavingEmptyPayload(t *testing.T) {
	req, err := http.NewRequest("PUT", "/", strings.NewReader(""))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-type", "application/json")

	ctx := context.WithValue(req.Context(), middleware.UserID, testUserID)
	req = req.WithContext(ctx)

	repoMock := &RepositoryMock{}
	h := NewHandler(repoMock)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.handlePut)

	handler.ServeHTTP(rr, req)

	expCode := http.StatusBadRequest
	if rr.Code != expCode {
		t.Errorf("expected status code %d but got %d", expCode, rr.Code)
	}
}

func TestGet(t *testing.T) {
	time_, err := time.Parse("January 2, 15:04:05, 2006", "January 2, 15:04:05, 2006")
	if err != nil {
		t.Fatal("Failed to parse date")
	}

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-type", "application/json")

	ctx := context.WithValue(req.Context(), middleware.UserID, testUserID)
	req = req.WithContext(ctx)

	repoMock := &RepositoryMock{}
	repoMock.FindFunc = func(string) ([]byte, time.Time, error) {
		return []byte("hello world"), time_, nil
	}
	h := NewHandler(repoMock)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.handleRead)

	handler.ServeHTTP(rr, req)

	expCode := http.StatusOK
	if rr.Code != expCode {
		t.Errorf("expected status code %d but got %d", expCode, rr.Code)
	}

	exp := `{"payload":"hello world","meta":{"modifiedAt":"2006-01-02T15:04:05Z"}}` + "\n"
	if rr.Body.String() != exp {
		t.Fatalf("expected %q, received %q", exp, rr.Body.String())
	}
}

func TestGetReturns204WhenNoContent(t *testing.T) {
	time_, err := time.Parse("January 2, 15:04:05, 2006", "January 2, 15:04:05, 2006")
	if err != nil {
		t.Fatal("Failed to parse date")
	}

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-type", "application/json")

	ctx := context.WithValue(req.Context(), middleware.UserID, testUserID)
	req = req.WithContext(ctx)

	repoMock := &RepositoryMock{}
	repoMock.FindFunc = func(string) ([]byte, time.Time, error) {
		return []byte(""), time_, ErrNoRows
	}
	h := NewHandler(repoMock)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.handleRead)

	handler.ServeHTTP(rr, req)

	expCode := http.StatusNoContent
	if rr.Code != expCode {
		t.Errorf("expected status code %d but got %d", expCode, rr.Code)
	}
}

func TestGetReturns500WhenRepositoryError(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-type", "application/json")

	ctx := context.WithValue(req.Context(), middleware.UserID, testUserID)
	req = req.WithContext(ctx)

	repoMock := &RepositoryMock{}
	repoMock.FindFunc = func(string) ([]byte, time.Time, error) {
		return []byte(""), time.Now(), errors.New("i failed")
	}
	h := NewHandler(repoMock)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.handleRead)

	handler.ServeHTTP(rr, req)

	expCode := http.StatusInternalServerError
	if rr.Code != expCode {
		t.Errorf("expected status code %d but got %d", expCode, rr.Code)
	}
}

func TestDelete(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-type", "application/json")

	ctx := context.WithValue(req.Context(), middleware.UserID, testUserID)
	req = req.WithContext(ctx)

	repoMock := &RepositoryMock{}
	repoMock.DeleteFunc = func(s string) error {
		return nil
	}
	h := NewHandler(repoMock)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.handleDelete)

	handler.ServeHTTP(rr, req)

	expCode := http.StatusOK
	if rr.Code != expCode {
		t.Errorf("expected status code %d but got %d", expCode, rr.Code)
	}

	expectedBody := "\"Deleted\"\n"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rr.Body.String())
	}
}

func TestDeleteReturns500WhenRepositoryError(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-type", "application/json")

	ctx := context.WithValue(req.Context(), middleware.UserID, testUserID)
	req = req.WithContext(ctx)

	repoMock := &RepositoryMock{}
	repoMock.DeleteFunc = func(s string) error {
		return errors.New("i failed")
	}
	h := NewHandler(repoMock)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.handleDelete)

	handler.ServeHTTP(rr, req)

	expCode := http.StatusInternalServerError
	if rr.Code != expCode {
		t.Errorf("expected status code %d but got %d", expCode, rr.Code)
	}
}
