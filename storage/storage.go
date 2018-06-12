package storage

import (
	"errors"

	"github.com/retranslator-solution/retranslator_server/models"
)

var (
	NotFound = errors.New("entry not found")
	AlreadyExist = errors.New("entry already exist")
)

type Storage interface {
	Get(name string) (*models.Resource, error)
	UpdateOrCreate(resource *models.Resource) error
	GetResourceNames() ([]string, error)
	Delete(name string) error
	Close() error
}
