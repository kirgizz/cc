package main

import (
	"app/auth"
	"app/migrations"
	"app/models"
	"app/server"
	"app/services"
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ivahaev/go-logger"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//createDataase()
//addUsers()
//addArticles()
//fmt.Print(models.GetArticlesByRating(0, "="))
//models.GetArticlesByUserNickname("john")

func OpenIndexHtml() string {
	b, err := ioutil.ReadFile("/home/evgeniy.sergeev/stuff/culture-city-golang/frontend/index.html")
	if err != nil {
		logger.Error(err)
	}
	return string(b)

}

const migrate = false

//LOGGING???
//ADD INTERFACES???

func main() {
	if migrate == true {
		migrations.CreateDBStruct()
	}

	srv := &http.Server{Addr: ":" + "localhost:8080", Handler: Routing()}

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt, os.Kill, syscall.SIGSTOP)

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			logger.Info("Listen and serve", err)
		}
	}()

	logger.Info("Server gracefully started at port 8080")
	<-stopChan // wait for SIGINT
	logger.Info("Shutting down server...")

	// shut down gracefully, but wait no longer than 5 seconds before halting
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)

	// close database connection
	services.GetInstanceDB().Close()
	logger.Info("Server gracefully stopped")

	logger.Info("Started")

}

func Routing() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	//r.HandleFunc("/api/login", server.Login).Methods("POST")

	r.Route("/api", func(r chi.Router) {
		r.Post("/login", auth.Login)
		r.With(auth.CheckSession).Post("/logout", auth.Logout)
		r.Post("/register", auth.Register)
		r.Get("/getArticles", models.GetArticles)
		r.With(auth.CheckSession).Post("/addArticle", models.AddArticle) // GET /articles
		r.With(auth.CheckSession).Post("/updateArticle", models.AddArticle)
		r.Get("/api/findArticleByName", server.NotImplemented)
		r.Get("/api/findArticle", server.NotImplemented)

	})
	return r
}
