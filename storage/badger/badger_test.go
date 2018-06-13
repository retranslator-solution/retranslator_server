package badger

import (
	"log"
	"os"
	"path"
	"reflect"
	"testing"

	"github.com/dgraph-io/badger"
	"github.com/retranslator-solution/retranslator_server/models"
	"github.com/retranslator-solution/retranslator_server/storage"
)

func TestStorageBadgerImpl(t *testing.T) {

	badgerTestDir := path.Join(os.TempDir(), "test_badger")

	opts := badger.DefaultOptions

	opts.Dir = badgerTestDir
	opts.ValueDir = badgerTestDir
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	defer os.RemoveAll(badgerTestDir)

	s := NewStorage(db)

	_, err = s.Get("unknown_key")

	if err != storage.NotFound {
		t.Fatalf("Expected `NotFound` error, got: %s", err)
	}
	newResource := models.Resource{
		Name: "test_0",
		String: []models.StringResource{
			{
				Name:  "sign_in_button",
				Value: "Sign-in",
			},
			{
				Name:  "logout_button",
				Value: "Logout",
			},
		},
		Array: []models.ArrayResource{
			{
				Name: "planets_array",
				Values: []string{
					"Mercury",
					"Venus",
					"Earth",
					"Mars",
				},
			},
		},
		Plural: []models.PluralResource{
			{
				Name: "number_of_songs",
				Values: []models.PluralValue{
					{
						Value:    "%d song found.",
						Quantity: "one",
					},
					{
						Value:    "%d songs found.",
						Quantity: "other",
					},
				},
			},
		},
	}

	err = s.Upsert(&newResource)
	if err != nil {
		t.Fatal()
	}

	resource, err := s.Get(newResource.Name)

	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(*resource, newResource) {
		t.Fatalf("resource and newResource should be equal")
	}

	newKeys := []string{
		newResource.Name,
		"test_1",
		"test_2",
		"test_3",
	}
	for _, k := range newKeys {
		newResource.Name = k
		err = s.Upsert(&newResource)
		if err != nil {
			t.Fatal()
		}
	}

	keys, err := s.GetResourceNames()
	if err != nil {
		t.Fatal(err)
	}

	for i, k := range keys {
		if newKeys[i] != k {
			t.Fatalf("Expected key (%s) on index (%d), got (%s)", newKeys[i], i, k)
		}
	}

	err = s.Delete(resource.Name)
	if err != nil {
		t.Fatal(err)
	}

	keys, err = s.GetResourceNames()
	if err != nil {
		t.Fatal(err)
	}

	for _, k := range keys {
		if k == resource.Name {
			t.Fatalf("Deleted key in resources list")
		}
	}

}
