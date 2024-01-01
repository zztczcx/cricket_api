package batting

import (
	"context"
	db "cricket/db/sqlc"
	"sync"
)

func (l *loader) Insert(playerChan <-chan player, wg *sync.WaitGroup, errors chan<- error) {
	defer wg.Done()

	for p := range playerChan {
		_, err := l.store.CreatePlayer(context.Background(), db.CreatePlayerParams(p))
		if err != nil {
                        errors <- err
		}
	}
}
