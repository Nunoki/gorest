package gorest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nunoki/go-util/pkg/array"
)

// ContentTypeMiddleware will verify both that the received request is of an application/json
// content-type, and that the client accepts application/json in return
func ContentTypeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if (c.Request.Method == "PUT" || c.Request.Method == "POST") &&
			!array.Contains(c.Request.Header["Content-Type"], "application/json") {
			c.AbortWithStatusJSON(
				http.StatusNotAcceptable,
				"Content-type expected to be application/json",
			)
			return
		}

		if !array.Contains(c.Request.Header["Accept"], "*/*") &&
			!array.Contains(c.Request.Header["Accept"], "application/json") {
			c.AbortWithStatusJSON(
				http.StatusNotAcceptable,
				"Client doesn't accept JSON response",
			)
			return
		}

		c.Next()
	}
}
