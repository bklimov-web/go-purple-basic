package files

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type JsonDb struct {
	filename string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		filename: name,
	}
}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.filename)

	if err != nil {
		fmt.Print("Не удалось создать файл")
	}

	defer file.Close()

	_, err = file.Write(content)

	if err != nil {
		fmt.Print("Не удалось записать файл")
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	content, err := os.ReadFile(db.filename)

	if err != nil {
		return nil, errors.New("FILE_READ_ERROR")
	}

	return content, nil
}

func ReadJson(name string) []byte {
	splitName := strings.Split(name, ".")
	extension := splitName[len(splitName)-1]

	if extension != "json" {
		fmt.Print("Это не json файл")
		return nil
	}

	content, err := os.ReadFile(name)

	if err != nil {
		fmt.Print("Не удалось прочитать файл")
		return nil
	}

	return content
}
