package main

import (
	"demo/password/bin"
	"demo/password/files"
	"demo/password/storage"
)

func main() {
	db := files.NewJsonDb("bins.json")
	binList := storage.GetBinStorage(db)

	binInstance := bin.NewBin("new bin", false)
	binList.AddBin(binInstance)
}
