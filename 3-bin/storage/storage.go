package storage

import "errors"

type StorageDb struct {
	url string
}

func NewStorageDb(name string) *StorageDb {
	return &StorageDb{
		url: name,
	}
}

func (db *StorageDb) Read() ([]byte, error) {
	return nil, errors.New("Not implemented jet")
}

func (db *StorageDb) Write(content []byte) error {
	return nil
}
