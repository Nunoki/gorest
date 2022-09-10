package gorest

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestDummyAuthOk(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(AuthMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer debug")
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

func TestDummyAuthForbidden(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(AuthMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(resp, req)

	res := resp.Result()

	if res.StatusCode != 401 {
		t.Fatalf("expected code 401, received %d", res.StatusCode)
	}
}
