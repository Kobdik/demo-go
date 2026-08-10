package main

import (
	"demo-go/bin/bins"
	"demo-go/bin/file"
	"demo-go/bin/storage"

	"fmt"
)

func main() {
	binList, err := bins.CreatetBinList(8)
	if err != nil {
		fmt.Println("Can't create BinList")
		return
	}
	binList.PrintFieldTag("Bins")

	a, err := bins.CreateBin(true, "Bin A")
	if err != nil {
		fmt.Println("Can't create private Bin")
	} else {
		binList.AddBin(*a)
	}

	b, err := bins.CreateBin(false, "Bin B")
	if err != nil {
		fmt.Println("Can't create public Bin")
	} else {
		binList.AddBin(*b)
	}

	fileName := "binList.json"
	ok, err := storage.WriteLocal(binList, fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !ok {
		fmt.Println("Не удалось сохранить список в json-файл")
		return
	}

	list, err := storage.ReadLocal(fileName)
	if err != nil {
		fmt.Printf("Can't read local file %s\n", fileName)
		return
	}
	fmt.Println("Local file read:")
	for ind, bin := range list.Bins {
		fmt.Println(ind, bin)
	}
	fmt.Printf("Bins updated at %v\n", list.UpdatedAt)

	fileName = "go.mod"
	data, err := file.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Can't read file %s cause %s\n", fileName, err)
		return
	}
	isItJson, err := file.IsItJson(data)
	if err != nil {
		fmt.Printf("Can't unmarshal %s\n", fileName)
		return
	}
	if isItJson {
		fmt.Printf("%s is json file\n", fileName)
		fmt.Println(string(data))
	} else {
		fmt.Printf("%s is not json file\n", fileName)
	}
}

func promptData(prompt string) string {
	fmt.Printf("%s: ", prompt)
	var res string
	fmt.Scan(&res)
	return res
}
