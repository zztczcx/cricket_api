package api

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
        "github.com/go-chi/chi/v5/middleware"
        "github.com/go-chi/jwtauth/v5"
)

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
        FileServer(s.router, "/docs", http.Dir("./docs/html"))
}


func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
