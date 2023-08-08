package db

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

var placeBucket = []byte("places")

// Don't select same place within the time period.
// ? What if user didn't add enough place to select different place for each day
var timePeriod = 5 * (24 * time.Hour)
var db *bolt.DB

// TODO: Add point logic
type Place struct {
	Id   int
	Name string
	Type string
	Last time.Time
}

func Init(path string) error {
	var err error

	db, err = bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return fmt.Errorf("init bolt db: %w", err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(placeBucket)
		return err
	})
	if err != nil {
		return fmt.Errorf("init bolt db: %w", err)
	}

	return nil
}

func AllPlaces() []Place {
	var places []Place

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(placeBucket)
		cur := b.Cursor()

		for key, value := cur.First(); key != nil; key, value = cur.Next() {
			places = append(places, deserializePlace(value))
		}

		return nil
	})

	return places
}

func AddPlace(place Place) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(placeBucket)
		id64, err := b.NextSequence()
		if err != nil {
			return fmt.Errorf("add place: %w", err)
		}
		id := int(id64)
		key := itob(id)
		place.Id = id

		value, err := serialize(place)
		if err != nil {
			return fmt.Errorf("add place: %w", err)
		}

		return b.Put(key, value)
	})

	panic("Implement add place")
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

func serialize(v any) ([]byte, error) {
	b, err := json.Marshal(v)
	return b, err
}

func deserializePlace(v []byte) Place {
	var place Place

	json.Unmarshal(v, &place)

	return place
}
