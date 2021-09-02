package kv_database

import (
	"io"
	"log"

	"github.com/cockroachdb/pebble"
)

type PebbleKV struct {
	db *pebble.DB
}

func (kv *PebbleKV) New() error {

	db, err := pebble.Open(
		"demo",
		&pebble.Options{},
	)

	if err != nil {
		return err
	}

	kv.db = db

	return nil
}

func (kv *PebbleKV) Get(key string) ([]byte, io.Closer, error) {

	value, closer, err := kv.db.Get([]byte(key))
	if err != nil {
		log.Println(err)
	}

	return value, closer, nil
}

func (kv *PebbleKV) Set(key []byte, value []byte) error {

	err := kv.db.Set(
		key,
		value,
		pebble.Sync,
	)
	if err != nil {
		log.Println(err)
	}

	return err
}
