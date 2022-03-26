package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQuery *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgres://postgres:password@localhost:5432/go_lang_app_db?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	testQuery = New(conn)
	os.Exit(m.Run())
}
