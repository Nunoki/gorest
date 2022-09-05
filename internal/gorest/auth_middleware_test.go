package gorest

import (
	"io"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestDummyAuthOk(t *testing.T) {
	ch := make(chan struct{})

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(AuthMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	go func(ch chan struct{}) {
		ch <- struct{}{}
		r.Run(":4200")
	}(ch)

	<-ch
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:4200", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer debug")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("expected code 200, received %d", resp.StatusCode)
	}
	content, _ := io.ReadAll(resp.Body)
	if string(content) != "ok" {
		t.Fatalf("expected body to equal \"ok\", received %s", string(content))
	}
}

func TestDummyAuthForbidden(t *testing.T) {
	ch := make(chan struct{})
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.Use(AuthMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	go func(ch chan struct{}) {
		ch <- struct{}{}
		r.Run(":4200")
	}(ch)

	<-ch
	resp, err := http.Get("http://localhost:4200")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 401 {
		t.Fatalf("expected code 401, received %d", resp.StatusCode)
	}
}
