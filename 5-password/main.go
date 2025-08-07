package main

import (
	"demo/password/account"
	"demo/password/encrypter"
	"demo/password/file"
	"fmt"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	enc := encrypter.NewEncrypter()
	db := file.NewJsonDb("data.vault")
	vault := account.GetVault(db, enc)

Menu:
	for {
		menuItem := getMenu()

		switch menuItem {
		case 1:
			createAccount(vault)
		case 2:
			findAccountByUrl(vault)
		case 3:
			findAccountByLogin(vault)
		case 4:
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}

func getMenu() int {
	return promptData([]string{
		"__Менеджер паролей__",
		"1. Создать аккаунт",
		"2. Найти аккаунт по url",
		"3. Найти аккаунт по логину",
		"4. Удалить аккаунт",
		"5. Выход",
		"Выберите вариант: ",
	})

}

func createAccount(vault *account.VaultWithDb) {
	login := getUserInput("логин")
	password := getUserInput("пароль")
	url := getUserInput("url")

	myAccount, err := account.NewAccount(login, password, url)

	if err != nil {
		fmt.Print(err)
		return
	}

	vault.AddAccount(myAccount)
}

// var findAccountMap := map[string]func

func findAccountByUrl(vault *account.VaultWithDb) {
	url := getUserInput("url")

	accounts := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})

	outputResult(&accounts)
}

func findAccountByLogin(vault *account.VaultWithDb) {
	login := getUserInput("логин")

	accounts := vault.FindAccounts(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})

	outputResult(&accounts)
}

func outputResult(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		fmt.Println("Аккаунт не найден")
		return
	}

	for _, acc := range *accounts {
		acc.OutputAccInfo()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := getUserInput("url для удаления")
	vault.DeleteAccountByUrl(url)
}

func getUserInput(data string) string {
	var userInput string
	fmt.Print("Введите ", data, ": ")
	fmt.Scanln(&userInput)
	return userInput
}

func promptData[T any](prompts []T) int {
	var res int

	for i, prompt := range prompts {
		if i < len(prompts)-1 {
			fmt.Println(prompt)
		} else {
			fmt.Print(prompt)
		}
	}
	fmt.Scan(&res)
	return res
}
