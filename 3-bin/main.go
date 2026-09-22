package main

import (
	"demo-go/bin/api"
	"demo-go/bin/bins"
	"demo-go/bin/config"
	"demo-go/bin/file"
	"demo-go/bin/storage"

	"fmt"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		fmt.Println("Can't create Config")
		return
	}
	api.SomeRequest(conf)

	fileName := "binList.json"
	// dependency injection by interface
	binList, err := bins.CreatetBinList(file.NewJsonDb(fileName), 8)
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
	// save list
	err = binList.Save()
	if err != nil {
		fmt.Println(err)
		return
	}
	// reload list
	err = binList.Load()
	if err != nil {
		fmt.Println("Can't load binList")
		return
	}
	fmt.Println("Loaded list:")
	for ind, bin := range binList.Bins {
		fmt.Println(ind, bin)
	}
	fmt.Printf("Bins updated at %v\n", binList.UpdatedAt)
	// dependency injection by interface
	binList2, err := bins.CreatetBinList(storage.NewStorageDb("someStorageName"), 8)
	binList2.AddBin(*a)
	binList2.AddBin(*b)
	for ind, bin := range binList2.Bins {
		fmt.Println(ind, bin)
	}
	fmt.Printf("Bins updated at %v\n", binList2.UpdatedAt)

	fileName = "go.mod"
	if file.IsItJsonExtension(fileName) {
		fmt.Printf("%s has json extension\n", fileName)
	} else {
		fmt.Printf("%s has not json extention\n", fileName)
	}
}

func promptData(prompt string) string {
	fmt.Printf("%s: ", prompt)
	var res string
	fmt.Scan(&res)
	return res
}
