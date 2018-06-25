package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/retranslator-solution/retranslator_server/application"
	"github.com/retranslator-solution/retranslator_server/server/handlers"
	log "github.com/sirupsen/logrus"
)

func RunServer(app *application.Application) {
	server := &http.Server{
		Addr:    ":8081",
		Handler: handlers.GetRouter(app),
	}

	go func() {
		log.Infoln("Server started")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of minute
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
