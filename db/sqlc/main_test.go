package db

import (
	"cricket/config"
	"log"
	"os"
	"testing"
)

var testStore Store

func TestMain(m *testing.M) {
        cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	connPool, err := NewDatabase(cfg.Database)
	if err != nil {
		panic(err)
	}

        testStore = NewStore(connPool)

        os.Exit(m.Run())
}
