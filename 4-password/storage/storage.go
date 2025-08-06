package storage

import (
	"demo/password/bin"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Db interface {
	Read() ([]byte, error)
	Write([]byte)
}

type BinStorage struct {
	BinList   []bin.Bin `json:"binList"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type BinStorageWithDb struct {
	BinStorage
	db Db
}

func (storage *BinStorageWithDb) ReadBinStorage() ([]byte, error) {
	content, err := storage.db.Read()

	if err != nil {
		return nil, errors.New("READ_ERROR")
	}

	return content, nil
}

func GetBinStorage(db Db) *BinStorageWithDb {
	binStorageContent, err := db.Read()

	if err != nil {
		return &BinStorageWithDb{
			BinStorage: BinStorage{
				BinList:   []bin.Bin{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}

	var binStorage BinStorage
	err = json.Unmarshal(binStorageContent, &binStorage)

	if err != nil {
		fmt.Println("Не удалось расшифровать файл, создан новый")
		return &BinStorageWithDb{
			BinStorage: BinStorage{
				BinList:   []bin.Bin{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}

	return &BinStorageWithDb{
		BinStorage: binStorage,
		db:         db,
	}
}

func (storage *BinStorageWithDb) writeStorage() error {
	storage.UpdatedAt = time.Now()

	storageJson, err := json.Marshal(storage.BinStorage)

	if err != nil {
		fmt.Print("Не удалось преобразовать")
		return errors.New("JSON_PARSE_ERROR")
	}

	storage.db.Write(storageJson)

	return nil
}

func (storage *BinStorageWithDb) AddBin(bin *bin.Bin) {
	storage.BinList = append(storage.BinList, *bin)

	err := storage.writeStorage()

	if err != nil {
		fmt.Print("Не удалось сохранить бин")
	}
}
