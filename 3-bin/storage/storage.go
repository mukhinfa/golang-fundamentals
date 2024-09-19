package storage

import (
	"encoding/json"
	"golang_fundamentals/3-bin/bins"
	"golang_fundamentals/3-bin/file"
)

func ReadBin(name string) (*bins.BinList, error) {
	file, err := file.ReadFile(name)
	if err != nil {
		return nil, err
	}
	var binList bins.BinList
	err = json.Unmarshal(file, &binList)
	return &binList, nil
}

func SaveBin(binList bins.BinList) {
	data, err := binList.ToBytes()
	if err != nil {
		return
	}
	file.WriteFile(data, "data.json")
}
