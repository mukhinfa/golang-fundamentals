package api

import (
	"golang_fundamentals/3-bin/config"

	"bytes"
	"fmt"
	"io"
	"net/http"
)

const genUrl = "https://api.jsonbin.io/v3/b"

type Api struct {
	config config.Config
}

func NewRequest() *Api {
	return &Api{
		config: *config.NewConfig(),
	}

}

func (c Api) Post(bin []byte) ([]byte, error) {
	baseUrl := genUrl
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, baseUrl, bytes.NewBuffer(bin))
	if err != nil {
		fmt.Println("Error creating HTTP request:")
	}
	return request(req, client, c)
}

func (c Api) Get(id string) ([]byte, error) {
	baseUrl := genUrl + "/" + id

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, baseUrl, nil)
	if err != nil {
		fmt.Println("Error creating HTTP request:")
	}
	return request(req, client, c)
}

func (c Api) Put(bin []byte, id string) ([]byte, error) {
	baseUrl := genUrl + "/" + id
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, baseUrl, bytes.NewBuffer(bin))
	if err != nil {
		fmt.Println("Error creating HTTP request:")
	}
	return request(req, client, c)
}

func (c Api) Delete(id string) ([]byte, error) {
	baseUrl := genUrl + "/" + id

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, baseUrl, nil)
	if err != nil {
		fmt.Println("Error creating HTTP request:")
	}
	return request(req, client, c)
}

func request(req *http.Request, client *http.Client, c Api) ([]byte, error) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Master-Key", c.config.Key)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error creating HTTP request:")
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа")
		return nil, err
	}
	return body, nil
}
