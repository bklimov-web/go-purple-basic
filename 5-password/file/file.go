package file

import (
	"errors"
	"fmt"
	"os"
)

type JsonDb struct {
	filename string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		filename: name,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	content, err := os.ReadFile(db.filename)

	if err != nil {
		fmt.Print("Не удалось прочитать файл")
		return nil, errors.New("READ_ERROR")
	}

	return content, nil
}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.filename)

	if err != nil {
		fmt.Print("Не удалось создать файл")
	}

	defer file.Close()

	_, err = file.Write(content)

	if err != nil {
		fmt.Print("Не удалось записать")
	}
}
