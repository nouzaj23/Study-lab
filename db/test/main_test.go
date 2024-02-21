package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"study_lab/config"
	db "study_lab/db/sqlc"
	"testing"
)

var testQueries *db.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	conf, err := config.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}
	testDB, err = sql.Open(conf.DBDriver, conf.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	testQueries = db.New(testDB)

	os.Exit(m.Run())
}
