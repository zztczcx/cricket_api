package db

import (
	"cricket/config"
	"database/sql"
	"errors"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	ensureUsingTestDB(cfg.Database.DatabaseURL)

	testDB, err = NewDatabase(cfg.Database)
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func TxStore(t *testing.T) Store {
	tx, err := testDB.Begin()
	if err != nil {
		t.Error(err)
	}

	t.Cleanup(func() {
		err := tx.Rollback()
		require.NoError(t, err)
	})

	return New(testDB).WithTx(tx)
}

func ensureUsingTestDB(url string) {
	if !strings.Contains(url, "test") {
		panic(errors.New("not using test database"))
	}
}
