package api

import (
	"cricket/config"
	db "cricket/db/sqlc"

	"github.com/go-chi/chi/v5"
)

// jwtutil -secret=secret -encode -claims='{"user_id":111}'
// hardcoded jwtToken for testing
const jwtToken = "BEARER eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMTF9._cLJn0xFS0Mdr_4L_8XF8-8tv7bHyOQJXyWaNsSqlEs"

func newTestServer(store db.Store) *Server {
        cfg := config.HTTPServer{
                JwtSecret: "secret", // TODO: using .test.env
        }

	srv := &Server{
		router: chi.NewRouter(),
		store:  store,
                cfg: cfg,
	}

	srv.routes()

	return srv
}
