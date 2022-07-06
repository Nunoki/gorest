package beetroot

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

const DefaultPayloadLimit = 1000 // #default_payload_limit

type Server struct {
	router *gin.Engine
}

// NewServer returns an instantiation of a Server with the repository from repo, and the
// AuthMiddleware attached, and sets up the route handlers
func NewServer(repo Repository) *Server {
	s := Server{
		router: gin.Default(),
	}

	// limit payload size to prevent large payload attack
	limit := getPayloadSizeLimit()
	s.router.Use(SizeLimitMiddleware(limit))

	// create handler group, so that we can extract /ping as a public route
	handler := NewHandler(repo)
	g := s.router.Group("")

	// auth middleware before routes, order matters because it sets the user id in the context
	publicKey := getJWTPublicKey()
	g.Use(AuthMiddleware(publicKey))

	// business logic goes here
	g.GET("/", handler.HandleRead)
	g.PUT("/", handler.HandleStore)
	g.DELETE("/", handler.HandleDelete)

	// ping pong
	s.router.GET("/ping", func(c *gin.Context) {
		c.String(
			http.StatusOK,
			"pong\n",
		)
	})

	return &s
}

// ServeHTTP just wraps Gin's ServeHTTP
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// TODO: move this and getPayloadSizeLimit() into main instead
// getJWTPublicKey returns the public key for verifying JWT's from Lettuce, which in its current
// iteration is passed via an environment variable
func getJWTPublicKey() string {
	pkey := os.Getenv("JWT_PUBLIC_KEY")
	if len(pkey) == 0 {
		log.Print("Missing env variable JWT_PUBLIC_KEY")
	}
	return pkey
}

// Returns the byte limit for the payload, which should be passed as an environment variable
// PAYLOAD_BYTE_LIMIT; if it isn't, then a default limit of DefaultPayloadLimit will be returned
func getPayloadSizeLimit() int64 {
	ls := os.Getenv("PAYLOAD_BYTE_LIMIT")
	limit, err := strconv.Atoi(ls)
	if err != nil || limit <= 0 {
		limit = DefaultPayloadLimit
		log.Printf(
			"Payload limit set to default of %d bytes (use environment variable PAYLOAD_BYTE_LIMIT to override)",
			limit,
		)
	} else {
		log.Printf("Payload limit set to %d bytes", limit)
	}

	return int64(limit)
}
