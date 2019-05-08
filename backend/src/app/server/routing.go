package server

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"net/http"
)

func (s Server) Routing() http.Handler {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3005/"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		//r.With(CheckSession).Post("/logout", s.NotImplemented)
		//r.With(CheckSession2).Get("/logout", s.NotImplemented)
		r.Get("/getArticle/{[0-9]+}", s.NotImplemented)
		r.Get("/getArticles", s.NotImplemented)
		//r.With(CheckSession).Post("/addArticle", s.NotImplemented)
		//r.With(CheckSession).Post("/updateArticle", s.NotImplemented)
		r.Post("/calculateRating", s.NotImplemented)
		r.Get("/findArticleByName", s.NotImplemented)
		r.Get("/findArticle", s.NotImplemented)
		r.Get("/deleteArticle/{[0-9]+}", s.NotImplemented)

		r.Get("/getPublications", s.NotImplemented)
		r.Get("/rubrics", s.getRubrics)
		r.Post("/register", s.registerUser)
		r.Post("/savePublication", s.savePublication)
		r.Post("/login", s.login)
		r.Post("/checkSession", s.checkSession)
		r.Post("/logout", s.logout)
	})

	return r
}


