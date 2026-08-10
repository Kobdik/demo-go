package bins

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
	"reflect"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"updatedAt"`
}

type BinList struct {
	Bins      []Bin `json:"bins"`
	UpdatedAt time.Time
}

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890@#$%")

func CreateBin(private bool, name string) (*Bin, error) {
	if name == "" {
		return nil, errors.New("не задано поле name")
	}
	return &Bin{
		Id:        generateId(16),
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}, nil
}

func (bin *Bin) ToBytes() ([]byte, error) {
	data, err := json.Marshal(bin)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func CreatetBinList(capacity int) (*BinList, error) {
	if capacity <= 0 {
		return nil, errors.New("задайте capacity > 0")
	}
	return &BinList{
		Bins:      make([]Bin, 0, capacity),
		UpdatedAt: time.Now(),
	}, nil
}

func (list *BinList) AddBin(bin Bin) {
	list.Bins = append(list.Bins, bin)
	list.UpdatedAt = time.Now()
}

func (list *BinList) ToBytes() ([]byte, error) {
	data, err := json.Marshal(list)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (list *BinList) PrintFieldTag(name string) {
	rt := reflect.TypeOf(list)
	sf, ok := rt.Elem().FieldByName(name)
	if !ok {
		fmt.Printf("Field by name %s not found\n", name)
	}
	fmt.Println(sf.Tag)
}

func generateId(n int) string {
	res := make([]rune, n)
	for i := range res {
		res[i] = runes[rand.IntN(n)]
	}
	return string(res)
}
