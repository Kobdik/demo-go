package main

import (
	"demo-go/bin/api"
	"demo-go/bin/bins"
	"demo-go/bin/file"
	"demo-go/bin/storage"

	"fmt"
)

func main() {
	binList, err := bins.CreateBinList(8)
	if err != nil {
		fmt.Println("Can't create BinList")
	} else {
		fmt.Println(binList)
	}
	err = api.GetBin()

	err = file.ReadFile()
	if err != nil {
		fmt.Println("Can't read file")
	}
	storage.WriteLocal()
}
