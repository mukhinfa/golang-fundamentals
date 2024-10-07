package storage

import (
	"encoding/json"
	"golang_fundamentals/3-bin/bins"
)

type StorageDb struct {
	name string
	db   FileDb
}

func newStorageDb(name string) *StorageDb {
	return &StorageDb{
		name: name,
	}
}

type FileDb interface {
	Read() ([]byte, error)
	Write(content []byte)
}

func (db *StorageDb) ReadBin(name string) (*bins.BinList, error) {
	file, err := db.db.Read()
	if err != nil {
		return nil, err
	}
	var binList bins.BinList
	err = json.Unmarshal(file, &binList)
	return &binList, nil
}

func (db *StorageDb) Save(binList bins.BinList) {
	data, err := binList.ToBytes()
	if err != nil {
		return
	}
	db.db.Write(data)
}
