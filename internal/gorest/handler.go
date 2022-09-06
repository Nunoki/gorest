package gorest

import (
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ErrInvalidJSON = errors.New("invalid JSON")

// MarshalJSON marshals the time into the standard Atom time format
func (t customTime) MarshalJSON() ([]byte, error) {
	return []byte(t.Time.UTC().Format("\"2006-01-02T15:04:05Z\"")), nil
}

// New returns a new instance of the Handler with the repo as the user repository
func NewHandler(repo Repository) Handler {
	return Handler{
		Repo: repo,
	}
}

// Returns the user ID as set up by the middleware after successful token verification. If the user
// ID variable has not been set up, it will panic.
func userIDFromAuth(c *gin.Context) string {
	userID, ok := c.Get("userID")
	if !ok {
		panic("couldn't read user ID from auth middleware")
	}
	return userID.(string)
}

// HandleRead tries to read the stored blob for the authenticated user, and then outputs it
// alongside some meta data in JSON format.
func (h Handler) HandleRead(c *gin.Context) {
	userID := userIDFromAuth(c)
	blob, modifiedAt, err := h.Repo.Find(userID)

	if err == ErrNoRows {
		c.AbortWithStatusJSON(
			http.StatusNoContent,
			"",
		)
		return
	}

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		Response{
			Payload: string(blob),
			Meta: ResponseMeta{
				ModifiedAt: customTime{Time: modifiedAt},
			},
		},
	)
}

// HandleStore stores the posted blob and ties it to the authenticated user. The blob will be
// validated by the repository implementation to be a valid JSON. It will output the number of bytes
// written.
func (h Handler) HandleStore(c *gin.Context) {
	userID := userIDFromAuth(c)
	b, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			err,
		)
		return
	}

	if len(b) == 0 {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			"Posted content is empty",
		)
		return
	}

	// JSON input is not being validated, because Postegres' JSON data type serves as the validator
	if err := h.Repo.Update(userID, b); err != nil {
		if err == ErrInvalidJSON {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				"Invalid JSON",
			)
			return
		}

		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	c.JSON(http.StatusOK, len(b))
}

// HandleDelete will delete any data tied to the authenticated user. It will output the appropriate
// messages depending on whether there was any data in there to begin with.
func (h Handler) HandleDelete(c *gin.Context) {
	userID := userIDFromAuth(c)
	err := h.Repo.Delete(userID)

	if err == ErrNoRows {
		c.AbortWithStatusJSON(
			http.StatusOK,
			"Nothing to delete",
		)
		return
	}

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			RespError{
				Message: err.Error(),
			},
		)
		return
	}

	// if we got here, the user was deleted
	c.JSON(http.StatusOK, "Deleted")
}
