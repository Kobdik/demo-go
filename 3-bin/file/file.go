package file

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func IsItJson(data []byte) (bool, error) {
	var v any
	err := json.Unmarshal(data, &v)
	if err != nil {
		return false, err
	}
	return true, nil
}

func WriteFile(name string, content string) error {
	file, err := os.Create(name)
	if err != nil {
		fmt.Printf("Can't create file %s\n", name)
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Printf("Can't write content %s\n", content)
		return err
	}
	return nil
}
