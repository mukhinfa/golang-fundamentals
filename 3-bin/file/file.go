package file

import (
	"fmt"

	"os"
	"strings"
)

func ReadFile(name string) ([]byte, error) {
	if !strings.HasSuffix(name, ".json") {
		error := fmt.Errorf("Файл должен быть \".json\"")
		return nil, error
	}
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.Write(content)
	if err != nil {
		file.Close()
		fmt.Println(err)
		return
	}
	defer file.Close()
}
