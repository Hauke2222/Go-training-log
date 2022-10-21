package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var liftsBucket = []byte("lifts")

var db *bolt.DB

type Lift struct {
	Key   int
	Value string
}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(liftsBucket)
		return err
	})

}

func CreateLift(lift string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(liftsBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(lift))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func AllLifts() ([]Lift, error) {
	var exercises []Lift
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(liftsBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			exercises = append(exercises, Lift{
				Key:   btoi(k),
				Value: string(v),
			})

		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return exercises, nil
}
func UpdateLift() {

}

func DeleteLift(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(liftsBucket)
		return b.Delete(itob(key))
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))

}
