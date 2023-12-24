package batting

import (
	"cricket/config"
	db "cricket/db/sqlc"
)

type Loader interface {
	Load() error
}

type loader struct {
	dataFile *string
	store  db.Store
}

func NewLoader(i *string) Loader {
        cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	connPool, err := db.NewDatabase(cfg.Database)
	if err != nil {
		panic(err)
	}

	return &loader{
		dataFile: i,
		store:  db.NewStore(connPool),
	}
}

func (l *loader) Load() error {
	dataSource := l.produce()
        done := make(chan struct{})
	l.startParser(dataSource, done)

        <-done
        return nil
}
