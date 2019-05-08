package main

import (
	"app/server"
	"app/services"
	"context"
	"flag"
	"github.com/ivahaev/go-logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	s server.Server
	migrate bool
)

func main() {
	flag.StringVar(&s.Port, "port", "8080", "port")
	flag.BoolVar(&migrate, "migrate", true, "migrate database via start")
	flag.Parse()

	if migrate == true {
		//migrations.CreateDBStruct()
		//migrations.UpdateDbStruct()
	}

	srv := &http.Server{Addr: ":" +s.Port, Handler: s.Routing()}

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt, os.Kill, syscall.SIGSTOP)


	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Error("Error while start server ", err)
			os.Exit(1)
		}
	}()

	logger.Info("Server gracefully started")
	<-stopChan // wait for SIGINT
	logger.Info("Shutting down server...")

	// shut down gracefully, but wait no longer than 5 seconds before halting
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)

	// close database connection
	services.GetInstanceDB().Close()
	logger.Info("Server gracefully stopped")

}
