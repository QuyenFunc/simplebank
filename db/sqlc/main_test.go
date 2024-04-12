package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/root?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

// tao ra bien testDB de luu ket qua sql
func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("connot connect to db")
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}
