package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"golang_fundamentals/3-bin/bins"
)

type Storage struct {
	name string
}

type StorageDb struct {
	Storage
	FileDb
}

func NewStorageDb(name string, db FileDb) *StorageDb {
	return &StorageDb{
		Storage: Storage{
			name: name,
		},
		FileDb: db,
	}
}

type FileDb interface {
	Read() ([]byte, error)
	Write(content []byte)
}

func (db *StorageDb) Read() (*bins.BinList, error) {
	file, err := os.ReadFile(db.name)
	if err != nil {
		fmt.Println("Ошибка чтения файла")
		return nil, err
	}
	var binList bins.BinList
	err = json.Unmarshal(file, &binList)
	if err != nil {
		fmt.Println("Ошибка анмаршаллинга")
		return nil, err
	}
	return &binList, nil
}

func (db *StorageDb) Save(binList bins.BinList) error {
	data, err := binList.ToBytes()
	if err != nil {
		fmt.Println("Ошибка перехода в байты")
		return err
	}
	db.Write(data)
	return nil
}
