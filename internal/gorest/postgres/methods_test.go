//go:build integration
// +build integration

package postgres

import (
	"bytes"
	"context"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/nunoki/gorest/internal/gorest"
)

var (
	ctx    = context.Background()
	client *Client
)

func TestMain(m *testing.M) {
	migrationDir := "../../../.docker/postgres/*.up.sql"

	connStr := ConnectionStringFromEnv()
	c, err := NewClient(ctx, connStr)
	if err != nil {
		log.Fatalln(err)
	}

	for _, q := range []string{
		`DROP SCHEMA IF EXISTS "test_schema" CASCADE;`,
		`CREATE SCHEMA "test_schema";`,
		`SET SEARCH_PATH TO "test_schema";`,
	} {
		if _, err := c.conn.Exec(context.Background(), q); err != nil {
			log.Fatalln(err)
		}
	}

	defer func() {
		query := `DROP SCHEMA "test_schema" CASCADE;`
		if _, err := c.conn.Exec(context.Background(), query); err != nil {
			log.Fatalln(err)
		}
	}()

	entries, err := filepath.Glob(migrationDir)
	if err != nil || len(entries) == 0 {
		log.Fatalln("migrations not found", err)
	}

	for _, file := range entries {
		migration, err := os.ReadFile(file)
		if err != nil {
			log.Fatalln(err)
		}

		_, err = c.conn.Exec(context.Background(), string(migration))
		if err != nil {
			log.Fatalln(err)
		}
	}

	client = c

	code := m.Run()
	os.Exit(code)
}

func TestFindNonexistentReturnsCorrectError(t *testing.T) {
	nonexistent := "f44fe12d-8bec-4720-845e-dbebcc053f90"
	if _, _, err := client.Find(nonexistent); err != gorest.ErrNoRows {
		t.Fatal("expected:", gorest.ErrNoRows, "received:", err)
	}
}

func TestDeleteNonexistentReturnsCorrectError(t *testing.T) {
	nonexistent := "f44fe12d-8bec-4720-845e-dbebcc053f91"
	if err := client.Delete(nonexistent); err != gorest.ErrNoRows {
		t.Fatal("expected:", gorest.ErrNoRows, "received:", err)
	}
}

func TestUpdateAndFetch(t *testing.T) {
	userID := "f44fe12d-8bec-4720-845e-dbebcc053f92"
	content := []byte(`{"number":1234653}`)
	start := time.Now()

	if err := client.Update(userID, content); err != nil {
		t.Fatal("should be able to update", err)
	}

	b, ts, err := client.Find(userID)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(b, content) {
		t.Fatalf("stored content does not match expected: %s received: %s", content, b)
	}

	if start.After(ts) {
		t.Fatalf(
			"returned modified time cannot be in the past, modified time: %s, current time: %s",
			ts,
			start,
		)
	}
}

func TestDeleteAndFetch(t *testing.T) {
	userID := "f44fe12d-8bec-4720-845e-dbebcc053f93"
	content := []byte("123")

	if err := client.Update(userID, content); err != nil {
		t.Fatal(err)
	}

	c, _, err := client.Find(userID)
	if err != nil {
		t.Fatal(err)
	}
	if c == nil {
		t.Fatal("failed to insert data for test", c)
	}

	if err := client.Delete(userID); err != nil {
		t.Fatal(err)
	}

	c, _, err = client.Find(userID)
	if err != gorest.ErrNoRows {
		t.Fatal("expected:", gorest.ErrNoRows, "received:", err)
	}
	if c != nil {
		t.Fatal("expected to be deleted, but is still present", c)
	}
}

func TestStoreComplexJSONAndFetch(t *testing.T) {
	userID := "f44fe12d-8bec-4720-845e-dbebcc053f94"
	jsonData := []byte(`{"name":"Mom","age":200,"isHappy":true,"children":[{"name":"Bob","favoriteToy":"teddy bear"},{"name":"Alice","favoriteToy":"teddy bear"}]}`)

	if err := client.Update(userID, jsonData); err != nil {
		t.Fatal(err)
	}

	d, _, err := client.Find(userID)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte(`{"name":"Mom","age":200,"isHappy":true,"children":[{"name":"Bob","favoriteToy":"teddy bear"},{"name":"Alice","favoriteToy":"teddy bear"}]}`)
	if !bytes.Equal(d, expected) {
		t.Fatal("expected:", string(expected), "received:", string(d))
	}
}
