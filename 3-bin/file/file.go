package file

import (
	"fmt"
	"os"
	"strings"
)

type JsonDb struct {
	filename string
}

func NewJsonDB(name string) *JsonDb {
	return &JsonDb{
		filename: name,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	if !strings.HasSuffix(db.filename, ".json") {
		error := fmt.Errorf("файл должен быть \".json\"")
		return nil, error
	}
	data, err := os.ReadFile(db.filename)
	if err != nil {
		_, err := os.Create(db.filename)
		if err != nil {
			fmt.Println("Ошибка создания файла")
			fmt.Println(err)
		}
	}
	return data, nil
}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		fmt.Println("Ошибка создания файла")
		fmt.Println(err)
	}
	_, err = file.Write(content)
	if err != nil {
		file.Close()
		fmt.Println("Ошибка записи в файл")
		fmt.Println(err)
		return
	}
	defer file.Close()
}
