package gorest

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware will call the appropriate authorization method, and then call the subsequent
// handler call if the authorization was successful. If it wasn't, a StatusUnauthorized status
// code will be output, and execution terminated
func AuthMiddleware(pubkey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := dummyAuth(c)
		// err := jwtAuth(c, pubkey)

		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				err.Error(),
			)
			return
		}
		c.Next()
	}
}

// dummyAuth is a debug/development authentication in which a dummy user id is being set, and
// authentication is marked as successful without doing any work
func dummyAuth(c *gin.Context) error {
	authHeader := c.GetHeader("authorization")

	if authHeader != "Bearer debug" {
		return errors.New("authentication failed, need bearer token of value \"debug\"")
	}

	c.Set("userID", "00000000-0000-0000-0000-000000000000")
	return nil
}
