package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"

	"github.com/daginwu/conscious_prog/struct_store/kv_database"
	"github.com/godruoyi/go-snowflake"
	"github.com/spf13/cast"
)

type User struct {
	IsEnabled bool
	ID        uint64
	IntVal    int
	FloatVal  float64
}

func NewUser() *User {

	id := snowflake.ID()

	return &User{
		IsEnabled: true,
		ID:        id,
		IntVal:    100,
		FloatVal:  0.1,
	}
}

type KV interface {
	New() error
	Get(key string) ([]byte, io.Closer, error)
	Set(key []byte, value []byte) error
}

func main() {

	// Declare KV interface
	var kv KV

	// Implementation KV
	kv = &kv_database.PebbleKV{}

	err := kv.New()
	if err != nil {
		log.Println(err)
	}

	// Init User struct
	wUser := NewUser()

	// Encode
	var wBuf bytes.Buffer
	enc := gob.NewEncoder(&wBuf)
	err = enc.Encode(wUser)
	if err != nil {
		log.Println(err)
	}

	// Pebble Set
	kv.Set(
		[]byte(cast.ToString(wUser.ID)),
		wBuf.Bytes(),
	)

	// Pebble Get
	var rUser User
	value, closer, err := kv.Get(cast.ToString(wUser.ID))

	// Decode
	var rBuf bytes.Buffer
	rBuf.Write(value)
	dec := gob.NewDecoder(&rBuf)
	err = dec.Decode(&rUser)
	if err != nil {
		log.Println(err)
	}
	// Relase
	closer.Close()

	// Print result
	fmt.Println(rUser)

}
