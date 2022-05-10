package beetroot

import (
	"errors"
	"time"
)

var ErrNoRows = errors.New("no rows")

//go:generate moq -out repository_moq_test.go . Repository
type Repository interface {
	Update(string, []byte) error
	Find(string) ([]byte, time.Time, error)
	Delete(string) error
}
