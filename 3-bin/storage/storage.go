package storage

import (
	"demo-go/bin/bins"
	"encoding/json"
	"os"
)

func WriteLocal(list *bins.BinList, name string) (bool, error) {
	data, err := json.Marshal(list)
	if err != nil {
		return false, err
	}
	file, err := os.Create(name)
	if err != nil {
		return false, err
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ReadLocal(name string) (list *bins.BinList, err error) {
	var data []byte
	data, err = os.ReadFile(name)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &list)
	return
}
