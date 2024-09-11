package main

import (
	"time"
)

type bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type binList struct {
	bins []bin
}

func createBin(id, name string, private bool) *bin {
	newBin := &bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
	return newBin
}

func createBinList(bin ...bin) *binList {
	newBinList := &binList{
		bins: bin,
	}
	return newBinList
}

func main() {

}
