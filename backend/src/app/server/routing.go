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
		r.Post("/login", Login)
		r.Get("/login", NotImplemented)
		r.Post("/checkSession", CheckSessionHandler)
		r.With(CheckSession).Post("/logout", Logout)
		r.With(CheckSession2).Get("/logout", Logout)
		r.Get("/getArticle/{[0-9]+}", GetArticle)
		r.Get("/getArticles", GetArticles)
		r.With(CheckSession).Post("/addArticle", AddArticle)
		r.With(CheckSession).Post("/updateArticle", AddArticle)
		r.Post("/calculateRating", CalculateRating)
		r.Get("/findArticleByName", NotImplemented)
		r.Get("/findArticle", NotImplemented)
		r.Get("/deleteArticle/{[0-9]+}", NotImplemented)

		r.Get("/getPublications", s.getPublications)
		r.Post("/register", s.registerUser)
	})

	return r
}


