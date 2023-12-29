package players

import (
	db "cricket/db/sqlc"

	"github.com/go-chi/chi/v5"
)

type playersHandler struct {
	store db.Store
}

func NewPlayersHandler(r chi.Router, store db.Store) {
	h := &playersHandler{store: store}
        r.Route("/players", func(r chi.Router) {
                r.Get("/most_runs", h.HandleMostRuns)
                r.Get("/active", h.HandleActive)
        })
}
