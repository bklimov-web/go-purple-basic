package main

import (
	"demo/password/api"
	"demo/password/bin"
	"demo/password/config"
	"demo/password/files"
	"demo/password/storage"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config := config.NewConfig()
	db := files.NewJsonDb("bins.json")
	binList := storage.GetBinStorage(db)
	apiInstance := api.NewApi(config)
	fmt.Print(apiInstance)

	binInstance := bin.NewBin("new bin", false)
	binList.AddBin(binInstance)
}
