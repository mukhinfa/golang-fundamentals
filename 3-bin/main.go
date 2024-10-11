package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"golang_fundamentals/3-bin/api"
	"golang_fundamentals/3-bin/bins"
	"golang_fundamentals/3-bin/file"
	"golang_fundamentals/3-bin/storage"
)

type responsePost struct {
	Metadata bins.Bin `json:"metadata"`
}

type responsePut struct {
	Metadata struct {
		ParentId string `json:"parentId"`
		Private  bool   `json:"private"`
	} `json:"metadata"`
}

type operationData struct {
	name        string
	id          string
	binFileName string
	FileDb
	HttpRequest
}

func newOperationData(name, id, binFileName string, fileDb FileDb, httpRequest HttpRequest) *operationData {
	return &operationData{
		name:        name,
		id:          id,
		binFileName: binFileName,
		FileDb:      fileDb,
		HttpRequest: httpRequest,
	}
}

type FileDb interface {
	Read() ([]byte, error)
	Write(content []byte)
}

type HttpRequest interface {
	Post([]byte) ([]byte, error)
	Get(string) ([]byte, error)
	Put([]byte, string) ([]byte, error)
	Delete(string) ([]byte, error)
}

func main() {
	var name, id, fileName string

	flag.StringVar(&name, "name", "", "Имя бина")
	flag.StringVar(&id, "id", "", "bin id")
	flag.StringVar(&fileName, "file", "", "Имя файла")

	create := flag.Bool("create", false, "create bin")
	update := flag.Bool("update", false, "update bin")
	delete := flag.Bool("delete", false, "delete bin")
	get := flag.Bool("get", false, "delete bin")
	list := flag.Bool("list", false, "bin list bin")

	flag.Parse()

	od := newOperationData(
		name,
		id,
		"binStorage/"+fileName,
		file.NewJsonDB("binList.json"),
		api.NewRequest(),
	)

	switch {
	case *create:
		od.createBin()
	case *update:
		od.updateBin()
	case *delete:
		od.deleteBin()
	case *get:
		od.getBin()
	case *list:
		od.listBin()
	}

}

func (od *operationData) createBin() {
	fmt.Println("Создание бина")

	data, err := od.Read()
	if err != nil {
		fmt.Println("Ошибка чтения файла:")
	}

	resp, err := od.Post(data)
	if err != nil {
		fmt.Println("Ошибка запроса")
	}

	var respData responsePost
	err = json.Unmarshal(resp, &respData)
	if err != nil {
		fmt.Println("Ошибка анмаршаллинга", err.Error())
	}

	bin := bins.CreateBin(respData.Metadata.Id, od.name, respData.Metadata.Private)
	storageDb := storage.NewStorageDb("binList.json", file.NewJsonDB("binList.json"))
	binList, err := storageDb.Read()
	if err != nil {
		fmt.Println("Ошибка чтения Бинлиста")
	}

	binList.Bins = append(binList.Bins, *bin)
	err = storageDb.Save(*binList)
	if err != nil {
		fmt.Println("Ошибка сохранения Бинлиста")
	}

	fmt.Println("Bin создан")
}

func (od *operationData) updateBin() {
	fmt.Println("Обновление бина")

	data, err := od.Read()
	if err != nil {
		panic(err.Error())
	}

	resp, err := od.Put(data, od.id)
	if err != nil {
		panic(err.Error())
	}

	var respData responsePut
	err = json.Unmarshal(resp, &respData)
	if err != nil {
		panic(err.Error())
	}

	storageDb := storage.NewStorageDb("binList.json", file.NewJsonDB("binList.json"))
	binList, err := storageDb.Read()
	if err != nil {
		panic(err.Error())
	}

	for i, v := range binList.Bins {
		if v.Id == od.id {
			updatedBin := bins.CreateBin(respData.Metadata.ParentId, v.Name, respData.Metadata.Private)
			arr := append(binList.Bins[:i], *updatedBin)
			arr = append(arr, binList.Bins[i+1:]...)
			binList.Bins = arr
			break
		}
	}
	err = storageDb.Save(*binList)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Bin обновлен")
}

func (od *operationData) deleteBin() {
	fmt.Println("Удаление бина")

	_, err := od.Delete(od.id)
	if err != nil {
		panic(err.Error())
	}

	storageDb := storage.NewStorageDb("binList.json", file.NewJsonDB("binList.json"))
	binList, err := storageDb.Read()
	if err != nil {
		panic(err.Error())
	}

	for i, v := range binList.Bins {
		if v.Id == od.id {
			binList.Bins = append(binList.Bins[:i], binList.Bins[i+1:]...)
			break
		}
	}

	err = storageDb.Save(*binList)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Bin удален")
}

func (od *operationData) getBin() {
	fmt.Println("Получение бина")

	resp, err := od.Get(od.id)
	if err != nil {
		panic(err.Error())
	}

	result, err := prettyString(resp)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Bin %s\n%s", od.id, result)
}

func (od *operationData) listBin() {
	fmt.Println("Список бинов")

	storageDb := storage.NewStorageDb("binList.json", file.NewJsonDB("binList.json"))
	binList, err := storageDb.Read()
	if err != nil {
		panic(err.Error())
	}

	for _, v := range binList.Bins {
		fmt.Printf("id %s name %s\n", v.Id, v.Name)
	}
}

func prettyString(str []byte) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, str, "", "  "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
