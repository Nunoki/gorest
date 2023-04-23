package gorest

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// // SizeLimitMiddleware will abort the request with a `413` response code if the request body is
// // larger than the limit number of bytes.
// func SizeLimitMiddleware(limit int64) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		if c.Request.ContentLength > limit {
// 			c.AbortWithStatusJSON(
// 				http.StatusRequestEntityTooLarge,
// 				"payload is too large",
// 			)
// 			return
// 		}

// 		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, limit)
// 		c.Next()
// 	}
// }
