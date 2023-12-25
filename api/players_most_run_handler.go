package api

import (
	db "cricket/db/sqlc"
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/render"
)

type mostRunResponse struct {
	Name string `json:"name"`
	Runs int64  `json:"runs"`
}

func newMostRunResponse(p db.Player) mostRunResponse {
	return mostRunResponse{
		Name: p.Name,
		Runs: p.Runs.Int64,
	}
}

// extra rendering for data
func (mr mostRunResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) handlePlayersMostRuns(w http.ResponseWriter, r *http.Request) {
	givenYear := r.URL.Query().Get("careerEndYear")

	var player db.Player
	var err error

	if givenYear != "" {
		player, err = s.store.GetPlayersOfMostRunsByCareerEndYear(r.Context(), db.ToNullInt64(givenYear))
	} else {
		player, err = s.store.GetPlayersOfMostRuns(r.Context())
	}

	if err != nil {
		if err == sql.ErrNoRows {
			render.Render(w, r, ErrNotFound)
			return
		}
		log.Printf("Error while querying most runs: %s\n", err)
		render.Render(w, r, ErrInternalServerError)
		return
	}

	mostRun := newMostRunResponse(player)
	render.Render(w, r, mostRun)
}
