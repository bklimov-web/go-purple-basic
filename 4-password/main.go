package main

import (
	"fmt"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList []Bin

func newBin(name string, private bool) *Bin {
	return &Bin{
		id:        "123",
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
}

func main() {
	bin := newBin("new bin", false)

	fmt.Print(bin)
}
