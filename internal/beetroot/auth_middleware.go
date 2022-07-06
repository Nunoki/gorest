package beetroot

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

/*
// jwtAuth will authenticate user based on the JWT in the authorization header
// TODO: Replace lettucejwt with generic JWT authentication
func jwtAuth(c *gin.Context, pubkey string) error {
	jwt := c.GetHeader("authorization")

	// REMOVE (used to make testing easier, because expired tokens are rejected)
	if jwt == "Bearer debug" {
		c.Set("userID", "f44fe12d-8bec-4720-845e-dbebcc053f9f")
		return nil
	}

	// TODO: Should lettucejwt.Read() do this part on its own?
	if len(jwt) > 7 && jwt[:7] == "Bearer " {
		jwt = jwt[7:]
	}

	claims, err := lettucejwt.Read(jwt, pubkey)

	if err != nil {
		return err
	}

	// REMOVE debug
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

	return nil
}
*/
