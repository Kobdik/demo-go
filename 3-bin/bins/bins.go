package bins

import (
	"errors"
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

func CreateBin(id string, private bool, name string) (*Bin, error) {
	if id == "" || name == "" {
		return nil, errors.New("не задан id или name")
	}
	return &Bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}, nil
}

func CreateBinList(capacity int) (*BinList, error) {
	if capacity <= 0 {
		return nil, errors.New("задайте capacity > 0")
	}
	return &BinList{
		bins: make([]Bin, 0, capacity),
	}, nil
}
