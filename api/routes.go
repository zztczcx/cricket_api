package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
        "github.com/go-chi/chi/v5/middleware"
        "github.com/go-chi/jwtauth/v5"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
}


func (s *Server) routes() {
        s.router.Use(middleware.Logger)
	s.router.Use(render.SetContentType(render.ContentTypeJSON))

	s.router.Get("/health", s.handleGetHealth)

        // Protected routes
        tokenAuth := jwtauth.New("HS256", []byte(s.cfg.JwtSecret), nil)
	s.router.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))

                r.Route("/api/v1/players", func(r chi.Router) {
                        r.Get("/most_runs", s.handlePlayersMostRuns)
                        r.Get("/active", s.handlePlayersActive)
                })
        })
}
