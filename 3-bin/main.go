package main

import "golang_fundamentals/3-bin/bins"

type Db interface {
	Read() (*bins.BinList, error)
	Save(binList bins.BinList)
}

func main() {

}
