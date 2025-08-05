package main

import (
	"demo/password/bin"
	"demo/password/storage"
)

func main() {
	binList := storage.GetBinStorage()

	binInstance := bin.NewBin("new bin", false)
	binList.AddBin(binInstance)
}
