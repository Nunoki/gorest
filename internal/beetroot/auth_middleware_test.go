package beetroot

import (
	"crypto/ed25519"
	"encoding/base64"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func TestValidJWT(t *testing.T) {
	log.SetOutput(ioutil.Discard) // REMOVE
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, jwt.RegisteredClaims{
		Issuer:    "Laravel",
		IssuedAt:  &jwt.NumericDate{time.Now().Add(-time.Second)}, // XXX: How to properly assign values here to get rid of the warning "composite literal uses unkeyed fields"?
		ExpiresAt: &jwt.NumericDate{time.Now().Add(time.Hour)},
		Subject:   "f44fe12d-8bec-4720-845e-dbebcc053f9e",
	})

	pubKey, privKey, _ := ed25519.GenerateKey(nil)
	tokenString, err := token.SignedString(privKey)
	if err != nil {
		t.Fatal("failed to generate signed token for test")
	}

	base64PubKey := base64.StdEncoding.EncodeToString(pubKey)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(AuthMiddleware(base64PubKey))

	handlerOk := false
	router.GET("/", func(c *gin.Context) {
		userID, _ := c.Get("userID")

		uid, ok := userID.(string)
		if !ok {
			t.Fatal("User id is not a string", uid)
		}
		if uid != "f44fe12d-8bec-4720-845e-dbebcc053f9e" {
			t.Fatal("User id does not match", uid)
		}

		body, _ := io.ReadAll(c.Request.Body)
		handlerOk = string(body) == `"hello world"`
	})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", strings.NewReader(`"hello world"`))
	r.Header.Set("Authorization", tokenString)
	router.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected response code %d, got %d", http.StatusOK, w.Code)
	}

	if !handlerOk {
		t.Fatal("Handler not called, or content doesn't match")
	}
}

func TestShouldFailWithInvalidJWT(t *testing.T) {
	log.SetOutput(ioutil.Discard) // REMOVE
	token := "invalid.token.string"

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(AuthMiddleware(""))
	router.GET("/", func(c *gin.Context) {
		t.Fatal("Handler method should not get called")
	})
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", strings.NewReader(`"hello world"`))
	r.Header.Set("Authorization", token)
	router.ServeHTTP(w, r)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("Expected code %d, received %d", http.StatusUnauthorized, w.Code)
	}
}
