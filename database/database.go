package database

import (
	"time"

	"go.etcd.io/bbolt"
)

type Database struct {
	client *bbolt.DB
}

func Open(path string) (*Database, error) {
	db, err := bbolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}

	d := &Database{client: db}
	if err := d.Set("groak", "started", time.Now().String()); err != nil {
		return nil, err
	}

	return d, nil
}

func (d *Database) Setup() error {
	return d.SaveSettings(&Settings{})
}

func (d *Database) Close() {
	d.client.Close()
}

func (d *Database) Get(bucket, key string) (string, error) {
	var value []byte
	err := d.client.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return nil
		}
		value = b.Get([]byte(key))
		return nil
	})
	return string(value), err
}

func (d *Database) Set(bucket, key, value string) error {
	return d.client.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}
		return b.Put([]byte(key), []byte(value))
	})
}
