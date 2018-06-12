package badger

import (
	"encoding/json"

	"github.com/dgraph-io/badger"
	"github.com/retranslator-solution/retranslator_server/models"
	"github.com/retranslator-solution/retranslator_server/storage"
)

type Storage struct {
	*badger.DB
}


func NewStorage(db *badger.DB) *Storage {
	return &Storage{
		DB: db,
	}
}

func (s *Storage) Get(name string) (*models.Resource, error) {
	var resource models.Resource

	err := s.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(name))
		if err != nil {
			return err
		}
		val, err := item.Value()
		if err != nil {
			return err
		}
		return json.Unmarshal(val, &resource)
	})

	if err == badger.ErrKeyNotFound {
		err = storage.NotFound
	}

	return &resource, err
}

func (s *Storage) UpdateOrCreate(resource *models.Resource) error {
	return s.Update(func(txn *badger.Txn) error {
		data, _ := json.Marshal(resource)
		return txn.Set([]byte(resource.Name), data)
	})
}

func (s *Storage) GetResourceNames() ([]string, error) {
	keys := make([]string, 0)
	err := s.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		for it.Rewind(); it.Valid(); it.Next() {
			keys = append(keys, string(it.Item().Key()))
		}
		it.Close()
		return nil
	})
	return keys, err
}

func (s *Storage) Delete(name string) error {
	return s.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(name))
	})
}
