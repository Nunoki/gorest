package gorest

import (
	"errors"
	"io"
	"net/http"

	"github.com/go-chi/render"
)

var ErrInvalidJSON = errors.New("invalid JSON")

// MarshalJSON marshals the time into the standard Atom time format
func (t customTime) MarshalJSON() ([]byte, error) {
	return []byte(t.Time.UTC().Format("\"2006-01-02T15:04:05Z\"")), nil
}

// NewHandler returns a new instance of the Handler with the repo as the user repository
func NewHandler(repo Repository) Handler {
	return Handler{
		Repo: repo,
	}
}

// userIDFromAuth returns the user ID from this request's context.
// If it couldn't be retreived, it panics.
func userIDFromAuth(r *http.Request) string {
	userID, ok := r.Context().Value(userID).(string)
	if !ok {
		panic("couldn't read user ID from auth middleware")
	}
	return userID
}

// handleRead tries to read and then output the stored blob for the authenticated user.
// If there is nothing stored, it will respond with 204.
func (h Handler) handleRead(w http.ResponseWriter, r *http.Request) {
	userID := userIDFromAuth(r)
	blob, modifiedAt, err := h.Repo.Find(userID)

	if err == ErrNoRows {
		render.Status(r, http.StatusNoContent)
		render.JSON(w, r, nil)
		return
	}

	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, RespError{err.Error()})
		return
	}

	resp := Response{
		Payload: string(blob),
		Meta: ResponseMeta{
			ModifiedAt: customTime{Time: modifiedAt},
		},
	}
	render.JSON(w, r, resp)
}

// handlePut stores the posted data into the database.
// The stored data will be associated with the authorized user.
func (h Handler) handlePut(w http.ResponseWriter, r *http.Request) {
	userID := userIDFromAuth(r)
	b, err := io.ReadAll(r.Body)

	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, RespError{err.Error()})
		return
	}

	if len(b) == 0 {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, RespError{"Posted content is empty"})
		return
	}

	// JSON input is not being validated, because Postegres' JSON data type serves as the validator
	if err := h.Repo.Update(userID, b); err != nil {
		if err == ErrInvalidJSON {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, RespError{"Invalid JSON"})
			return
		}

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, RespError{err.Error()})
		return
	}

	render.JSON(w, r, len(b))
}

// handleDelete will delete any stored data tied to the authenticated user
// The response message will indicate if there was nothing to delete.
// The response status will be 200 if something was deleted, or if it was already empty.
func (h Handler) handleDelete(w http.ResponseWriter, r *http.Request) {
	userID := userIDFromAuth(r)
	err := h.Repo.Delete(userID)

	if err == ErrNoRows {
		render.JSON(w, r, "Nothing to delete")
		return
	}

	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, RespError{err.Error()})
		return
	}

	// if we got here, the user was deleted
	render.JSON(w, r, "Deleted")
}
