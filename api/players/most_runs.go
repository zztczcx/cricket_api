package players

import (
	"cricket/api/errors"
	db "cricket/db/sqlc"
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/render"
        "github.com/guregu/null"
)

type mostRunResponse struct {
	Name string `json:"name"`
	Runs null.Int  `json:"runs"`
}

func newMostRunResponse(p db.Player) mostRunResponse {
	return mostRunResponse{
		Name: p.Name,
                Runs: null.Int{NullInt64: p.Runs},
	}
}

// extra rendering for data
func (mr mostRunResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *playersHandler) HandleMostRuns(w http.ResponseWriter, r *http.Request) {
	givenYear := r.URL.Query().Get("careerEndYear")

	var player db.Player
	var err error

	if givenYear != "" {
		player, err = h.store.GetPlayerOfMostRunsByCareerEndYear(r.Context(), db.ToNullInt64(givenYear))
	} else {
		player, err = h.store.GetPlayerOfMostRuns(r.Context())
	}

	if err != nil {
		if err == sql.ErrNoRows {
			render.Render(w, r, errors.ErrNotFound)
			return
		}
		log.Printf("Error while querying most runs: %s\n", err)
		render.Render(w, r, errors.ErrInternalServerError)
		return
	}

	mostRun := newMostRunResponse(player)
	render.Render(w, r, mostRun)
}
