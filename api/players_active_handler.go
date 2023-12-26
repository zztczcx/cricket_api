package api

import (
	db "cricket/db/sqlc"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
)

type activePlayersResponse struct {
	Names []string `json:"names"`
}

func newActivePlayersResponse(players []db.Player) activePlayersResponse {
	names := make([]string, len(players))

	for i, p := range players {
		names[i] = p.Name
	}

	return activePlayersResponse{
		Names: names,
	}
}

// extra rendering for data
func (mr activePlayersResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) handlePlayersActive(w http.ResponseWriter, r *http.Request) {
	givenYear := r.URL.Query().Get("careerYear")
	if givenYear == "" {
		err := &ErrResponse{
			HTTPStatusCode: 400,
			StatusText:     "Bad request",
			ErrorText:      "careerYear is required as query parameter",
		}

		render.Render(w, r, err)
		return
	}

	year, err := strconv.Atoi(givenYear)
	if err != nil {
		err := &ErrResponse{
			HTTPStatusCode: 400,
			StatusText:     "Bad request",
			ErrorText:      "careerYear should be a number",
		}
		render.Render(w, r, err)
		return
	}

	p := db.GetPlayersByCareerYearParams{
		CareerYear: sql.NullInt64{Int64: int64(year), Valid: true},
	}

	players, err := s.store.GetPlayersByCareerYear(r.Context(), p)
	if err != nil {
		if err == sql.ErrNoRows {
			render.Render(w, r, ErrNotFound)
			return
		}
		log.Printf("Error while querying active players: %s\n", err)
		render.Render(w, r, ErrInternalServerError)
		return
	}

	playersResp := newActivePlayersResponse(players)
	render.Render(w, r, playersResp)

}
