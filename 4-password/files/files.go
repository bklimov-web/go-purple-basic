package files

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// func ReadFile() {

// }

func ReadJson(name string) {
	splitName := strings.Split(name, ".")
	extension := splitName[len(splitName)-1]

	if extension != "json" {
		fmt.Print("Это не json файл")
		return
	}

	content, err := os.ReadFile(name)

	if err != nil {
		fmt.Print("Не удалось прочитать файл")
		return
	}

	fmt.Println(content)
}

func WriteFile(name string, content []byte) error {
	file, err := os.Create(name)

	if err != nil {
		fmt.Print("Не удалось создать файл")
		return errors.New("FILE_CREATE_ERROR")
	}

	defer file.Close()

	_, err = file.Write(content)

	if err != nil {
		fmt.Print("Не удалось записать файл")
		return errors.New("FILE_WRITE_ERROR")
	}

	return nil
}
