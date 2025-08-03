package bin

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

func NewBin(name string, private bool) *Bin {
	return &Bin{
		id:        "123",
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
}

func NewBinList() *BinList {
	return &BinList{}
}

func main() {
	bin := NewBin("new bin", false)
	binList := NewBinList()

	fmt.Print(bin)
	fmt.Print(binList)
}
