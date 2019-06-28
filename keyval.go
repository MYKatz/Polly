package main

import (
	"log"

	"github.com/dgraph-io/badger"
)

type KeyVal struct {
	db *badger.DB
}

func NewKeyVal() *KeyVal {
	opts := badger.DefaultOptions
	opts.Dir = "/tmp/badger"
	opts.ValueDir = "/tmp/badger"
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()
	return &KeyVal{db}
}

func (kv *KeyVal) Set(key string, val string) error {
	err := kv.db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), []byte(val))
		return err
	})
	return err
}

func (kv *KeyVal) Get(key string) (string, error) {
	var valCopy []byte
	err := kv.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		valCopy, err = item.Value()

		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	return string(valCopy), err
}
