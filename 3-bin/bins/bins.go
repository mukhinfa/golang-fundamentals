package bins

import (
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
