package bin

import (
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createAt"`
	Name      string    `json:"name"`
}

func NewBin(name string, private bool) *Bin {
	return &Bin{
		Id:        "123",
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}
