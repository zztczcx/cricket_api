package batting

import (
	"context"
	db "cricket/db/sqlc"
	"log"
	"sync"
)

func (l *loader) Insert(playerChan <-chan player, wg *sync.WaitGroup) {
        defer wg.Done()

        for p := range playerChan {
                _, err := l.store.CreatePlayer(context.Background(), db.CreatePlayerParams(p))
                if err != nil {
                        log.Println(err)
                }
        }
}
