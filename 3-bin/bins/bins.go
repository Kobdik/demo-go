package bins

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
	"reflect"
	"time"
)

type Db interface {
	Read() ([]byte, error)
	Write([]byte) error
}

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

type BinList struct {
	Bins      []Bin `json:"bins"`
	UpdatedAt time.Time
}

type BinListWithDb struct {
	BinList
	db Db
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

func CreatetBinList(db Db, capacity int) (*BinListWithDb, error) {
	if capacity <= 0 {
		return nil, errors.New("задайте capacity > 0")
	}
	content, err := db.Read()
	if err != nil {
		return &BinListWithDb{
			BinList: BinList{
				Bins:      make([]Bin, 0, capacity),
				UpdatedAt: time.Now(),
			},
			db: db,
		}, nil
	}
	var binList BinListWithDb = BinListWithDb{db: db}
	err = json.Unmarshal(content, &binList)
	if err != nil {
		return nil, errors.New("Не удалось распарсить сохраненный json")
	}
	return &binList, nil
}

func (list *BinListWithDb) AddBin(bin Bin) {
	list.Bins = append(list.Bins, bin)
	list.UpdatedAt = time.Now()
}

func (list *BinListWithDb) ToBytes() ([]byte, error) {
	data, err := json.Marshal(list.BinList)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (list *BinListWithDb) PrintFieldTag(name string) {
	rt := reflect.TypeOf(list)
	sf, ok := rt.Elem().FieldByName(name)
	if !ok {
		fmt.Printf("Field by name %s not found\n", name)
	}
	fmt.Println(sf.Tag)
}

func (list *BinListWithDb) Load() error {
	data, err := list.db.Read()
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &list.BinList)
	return err
}

func (list *BinListWithDb) Save() error {
	list.UpdatedAt = time.Now()
	data, err := list.ToBytes()
	if err != nil {
		return err
	}
	return list.db.Write(data)
}

func generateId(n int) string {
	res := make([]rune, n)
	for i := range res {
		res[i] = runes[rand.IntN(n)]
	}
	return string(res)
}
