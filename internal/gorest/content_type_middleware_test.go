package gorest

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	r = gin.New()
	r.Use(ContentTypeMiddleware())

	r.PUT("/put", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	r.GET("/get", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	m.Run()
}

func TestContentTypeMiddlewareOk(t *testing.T) {
	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/put", nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(resp, req)

	res := resp.Result()
	if res.StatusCode != 200 {
		t.Fatalf("expected code 200, received %d", res.StatusCode)
	}
	content, _ := io.ReadAll(res.Body)
	if string(content) != "ok" {
		t.Fatalf("expected body to equal \"ok\", received %s", string(content))
	}
}

func TestContentTypeMiddlewareFailsWhenNoAcceptHeader(t *testing.T) {
	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/put", nil)
	// req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(resp, req)

	res := resp.Result()
	if res.StatusCode != 406 {
		t.Fatalf("expected code 406, received %d", res.StatusCode)
	}
}

func TestContentTypeMiddlewareFailsWhenNoContentTypeHeader(t *testing.T) {
	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/put", nil)
	req.Header.Set("Accept", "application/json")
	// req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(resp, req)

	res := resp.Result()
	if res.StatusCode != 406 {
		t.Fatalf("expected code 406, received %d", res.StatusCode)
	}
}

func TestContentTypeMiddlewareContentTypeNotRequiredWhenGETMethod(t *testing.T) {
	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/get", nil)
	req.Header.Set("Accept", "application/json")
	// req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(resp, req)

	res := resp.Result()
	if res.StatusCode != 200 {
		t.Fatalf("expected code 200, received %d", res.StatusCode)
	}
	content, _ := io.ReadAll(res.Body)
	if string(content) != "ok" {
		t.Fatalf("expected body to equal \"ok\", received %s", string(content))
	}
}
