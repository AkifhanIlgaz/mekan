package db

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

var placeBucket = []byte("places")

// Don't select same place within the time period.
// ? What if user didn't add enough place to select different place for each day
var DefaultTimePeriod = 14 * (24 * time.Hour)

type PlaceService struct {
	DB         *bolt.DB
	TimePeriod time.Duration
}

type Place struct {
	Name string
	Type string
	Last time.Time
}

func Init(path string, timePeriod ...time.Duration) (*PlaceService, error) {
	placeService := PlaceService{
		TimePeriod: DefaultTimePeriod,
	}
	if len(timePeriod) > 0 {
		placeService.TimePeriod = timePeriod[0]
	}

	db, err := bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, fmt.Errorf("init bolt db: %w", err)
	}
	placeService.DB = db

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(placeBucket)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("init bolt db: %w", err)
	}

	return &placeService, nil
}
