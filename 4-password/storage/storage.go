package storage

import (
	"demo/password/bin"
	"demo/password/files"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type BinStorage struct {
	BinList   []bin.Bin `json:"binList"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func ReadBinStorage() {
	files.ReadJson("bins.json")
}

func GetBinStorage() *BinStorage {
	binStorageContent, err := os.ReadFile("bins.json")

	if err != nil {
		return &BinStorage{
			BinList:   []bin.Bin{},
			UpdatedAt: time.Now(),
		}
	}

	var binStorage BinStorage
	err = json.Unmarshal(binStorageContent, &binStorage)

	if err != nil {
		fmt.Println("Не удалось расшифровать файл, создан новый")
		return &BinStorage{
			BinList:   []bin.Bin{},
			UpdatedAt: time.Now(),
		}
	}

	return &binStorage
}

func (storage *BinStorage) writeStorage() error {
	storage.UpdatedAt = time.Now()

	storageJson, err := json.Marshal(storage)

	if err != nil {
		fmt.Print("Не удалось преобразовать")
		return errors.New("JSON_PARSE_ERROR")
	}

	files.WriteFile("bins.json", storageJson)

	return nil
}

func (storage *BinStorage) AddBin(bin *bin.Bin) {
	storage.BinList = append(storage.BinList, *bin)

	err := storage.writeStorage()

	if err != nil {
		fmt.Print("Не удалось сохранить бин")
	}
}
