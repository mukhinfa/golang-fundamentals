package bins

import (
	"encoding/json"
	"fmt"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

type BinList struct {
	Bins []Bin `json:"bins"`
}

func CreateBin(id, name string, private bool) *Bin {
	newBin := &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
	return newBin
}

func CreateBinList(bin ...Bin) *BinList {
	newBinList := &BinList{
		Bins: bin,
	}
	return newBinList
}

func (binList *BinList) ToBytes() ([]byte, error) {
	file, err := json.Marshal(binList)
	if err != nil {
		fmt.Println("Ошибка маршаллинга")
		return nil, err
	}
	return file, nil
}
