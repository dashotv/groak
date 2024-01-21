package database

import (
	"errors"
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
	if db == nil {
		return nil, errors.New("failed to open database")
	}

	d := &Database{client: db}
	if err := d.Set("groak", "started", time.Now().String()); err != nil {
		return nil, err
	}

	return d, nil
}

func (d *Database) Setup() error {
	settings, err := d.GetSettings()
	if err != nil {
		return err
	}
	return d.SaveSettings(settings)
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

func (d *Database) Delete(bucket, key string) error {
	return d.client.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return nil
		}
		return b.Delete([]byte(key))
	})
}

func (d *Database) List(bucket string) ([]string, error) {
	var keys []string
	err := d.client.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return nil
		}
		return b.ForEach(func(k, v []byte) error {
			keys = append(keys, string(k))
			return nil
		})
	})
	return keys, err
}

func (d *Database) ListBuckets() ([]string, error) {
	var buckets []string
	err := d.client.View(func(tx *bbolt.Tx) error {
		return tx.ForEach(func(name []byte, b *bbolt.Bucket) error {
			buckets = append(buckets, string(name))
			return nil
		})
	})
	return buckets, err
}

func (d *Database) DeleteBucket(bucket string) error {
	return d.client.Update(func(tx *bbolt.Tx) error {
		return tx.DeleteBucket([]byte(bucket))
	})
}
