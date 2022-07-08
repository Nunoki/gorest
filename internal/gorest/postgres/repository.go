package postgres

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/nunoki/gorest/internal/gorest"
)

const queryFind = `
	SELECT
		"content",
		CASE
			WHEN "modified_at" IS NULL
			THEN "created_at"
			ELSE "modified_at"
		END
	FROM   "json_blob"
	WHERE  user_id = $1
	LIMIT  1
	`
const queryUpdate = `
	INSERT INTO "json_blob"("user_id", "content")
	VALUES($1, $2)
	ON CONFLICT("user_id") DO UPDATE SET "content" = $2
	`
const queryDelete = `DELETE FROM "json_blob" WHERE user_id = $1`

// Find will try to load the stored blob for the specified userID, and return its contents
// and the time it was last modified.
func (r Client) Find(userID string) ([]byte, time.Time, error) {
	var blob json.RawMessage
	var modifiedAt time.Time

	err := r.conn.
		QueryRow(context.Background(), queryFind, userID).
		Scan(
			&blob,
			&modifiedAt,
		)

	if err == pgx.ErrNoRows {
		return blob, modifiedAt, gorest.ErrNoRows
	}

	if err != nil {
		return blob, modifiedAt, err
	}

	return blob, modifiedAt, nil
}

// Update will update the stored blob for the specified userID, as long as the provided blob is a
// valid JSON; if the validation fails, a special gorest.ErrInvalidJSON error is returned.
func (r Client) Update(userID string, content []byte) error {
	_, err := r.conn.Exec(
		context.Background(),
		queryUpdate,
		userID,
		content,
	)

	// XXX: Is there a more elegant way to check for this type of error?
	if err != nil && strings.Contains(err.Error(), "22P02") {
		return err
		// TODO: this condition is also satisfied when the userID is in an incorrect format!
		// return gorest.ErrInvalidJSON
	}

	return err
}

// Delete will delete any stored blob tied to the specified userID. The special error
// gorest.ErrNoRows is returned if there was no data to delete for the user.
func (r Client) Delete(userID string) error {
	tag, err := r.conn.Exec(
		context.Background(),
		queryDelete,
		userID,
	)

	if err != nil {
		return err
	}

	if tag.RowsAffected() < 1 {
		return gorest.ErrNoRows
	}

	return nil
}
