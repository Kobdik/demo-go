package file

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type JsonDb struct {
	filename string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		filename: name,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.filename)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func IsItJsonContent(data []byte) (bool, error) {
	var v any
	err := json.Unmarshal(data, &v)
	if err != nil {
		return false, err
	}
	return true, nil
}

func IsItJsonExtension(name string) bool {
	if ext := filepath.Ext(name); ext != ".json" {
		return false
	}
	return true
}

func (db *JsonDb) Write(content []byte) error {
	file, err := os.Create(db.filename)
	if err != nil {
		fmt.Printf("Can't create file %s\n", db.filename)
		return err
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		fmt.Printf("Can't write content to file %s\n", db.filename)
		return err
	}
	return nil
}
