package db

import (
	"database/sql"
	"github.com/baguss42/simplebank/utils"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	testDB      *sql.DB
	testQueries *Queries
)

func TestMain(m *testing.M) {
	var err error

	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load configuration", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the database: ", err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}
