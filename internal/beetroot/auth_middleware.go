package beetroot

import (
	"log"
	"net/http"
	"time"

	"github.com/binogi/go-pkg/lettucejwt"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware will call the Lettuce JWT verifier in order to confirm that the requester is an
// authenticated user as identified by Lettuce. It will then set an app variable containing the
// Subject from the JWT as `userID` and call the next handler. If authorization fails, the request
// is aborted with 401 and no handlers are being called.
// It will use the public key pubkey for token verification
func AuthMiddleware(pubkey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt := c.GetHeader("authorization")

		// REMOVE (used to make testing easier, because expired tokens are rejected)
		if jwt == "Bearer debug" {
			c.Set("userID", "f44fe12d-8bec-4720-845e-dbebcc053f9f")
			c.Next()
			return
		}

		// TODO: Should lettucejwt.Read() do this part on its own?
		if len(jwt) > 7 && jwt[:7] == "Bearer " {
			jwt = jwt[7:]
		}

		claims, err := lettucejwt.Read(jwt, pubkey)

		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				err.Error(),
			)
			return
		}

		// TODO: remove
		expiration := time.Unix(claims.ExpiresAt, 0).UTC()
		log.Printf(
			"sub: %s, iat: %d, exp: %d, iss: %s, expires at %s\n",
			claims.Subject,
			claims.IssuedAt,
			claims.ExpiresAt,
			claims.Issuer,
			expiration.Format(time.RFC1123Z),
		)

		c.Set("userID", claims.Subject)
		c.Next()
	}
}
