package main

import (
	"fmt"
	"json-cli/bins"
	"json-cli/storage"
)

func main() {

	storageHandler := storage.JSONStorage{}

	binList := &bins.BinList{}
	bin := bins.NewBin("1", false, "Test bin")
	binList.AddBin(bin)

	filePath := "bins.json"
	err := storageHandler.Save(filePath, binList)
	if err != nil {
		fmt.Println("ERROR SAVE: ", err)
		return
	}

	loadedBinList, err := storageHandler.Load(filePath)
	if err != nil {
		fmt.Println("ERROR LOAD: ", err)
		return
	}
	fmt.Printf("LOAD DATA: %+v\n", loadedBinList)

}
