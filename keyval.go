package main

import (
	"log"

	"github.com/dgraph-io/badger"
)

type KeyVal struct {
	test string
}

func OpenDB() {
	opts := badger.DefaultOptions("/tmp/badger")
	opts.Dir = "/tmp/badger"
	opts.ValueDir = "/tmp/badger"
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
