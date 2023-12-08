package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	sqlc "github.com/0xGeN02/UDEMY/GOLang/simpleBank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *sqlc.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = sqlc.New(testDB)
	os.Exit(m.Run())
}
