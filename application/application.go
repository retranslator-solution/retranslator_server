package application

import (
	"github.com/dgraph-io/badger"
	"github.com/retranslator-solution/retranslator_server/configs"
	"github.com/retranslator-solution/retranslator_server/storage"
	badgerStorage "github.com/retranslator-solution/retranslator_server/storage/badger"
	log "github.com/sirupsen/logrus"
)

type Application struct {
	Storage storage.Storage
}

func NewApplication(config configs.Config) *Application {
	var app Application

	switch config.Storage.Backend {
	case configs.BadgerBackend:
		opts := badger.DefaultOptions
		opts.Dir = config.Storage.Badger.Path
		opts.ValueDir = config.Storage.Badger.Path
		db, err := badger.Open(opts)
		if err != nil {
			log.Fatal(err)
		}

		app.Storage = badgerStorage.NewStorage(db)
	default:
		log.Fatalf("storage backend '%s' not supported", config.Storage.Backend)
	}

	return &app
}
