package bins

import (
	"encoding/json"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	bins []Bin
}

func CreateBin(id, name string, private bool) *Bin {
	newBin := &Bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
	return newBin
}

func CreateBinList(bin ...Bin) *BinList {
	newBinList := &BinList{
		bins: bin,
	}
	return newBinList
}

func (binList *BinList) ToBytes() ([]byte, error) {
	file, err := json.Marshal(binList)
	if err != nil {
		return nil, err
	}
	return file, nil
}
