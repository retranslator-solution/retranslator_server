package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/dgraph-io/badger"
	_ "github.com/lib/pq"
	"github.com/retranslator-solution/retranslator_server/application"
	"github.com/retranslator-solution/retranslator_server/server/handlers"
	badgerStore "github.com/retranslator-solution/retranslator_server/storage/badger"
	log "github.com/sirupsen/logrus"
)

func RunServer() {

	// todo: use configs
	badgerDir := path.Join(os.TempDir(), "test_badger")

	opts := badger.DefaultOptions

	opts.Dir = badgerDir
	opts.ValueDir = badgerDir
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	storage := badgerStore.NewStorage(db)
	app := &application.Application{
		Storage: storage,
	}

	server := &http.Server{
		Addr:    ":8080",
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
