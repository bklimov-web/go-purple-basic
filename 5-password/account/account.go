package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewAccount(login, password, urlString string) (*Account, error) {
	_, err := url.ParseRequestURI(urlString)

	if login == "" {
		return nil, errors.New("LOGIN_ERROR")
	}

	if err != nil {
		return nil, errors.New("URL_ERROR")
	}

	myAccount := &Account{
		Login:     login,
		Password:  password,
		Url:       urlString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if myAccount.Password == "" {
		myAccount.generatePassword(10)
	}

	return myAccount, nil
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)

	allowedSymbols := []rune("Привет!)")

	for i := range res {
		index := rand.IntN(len(allowedSymbols))

		res[i] = allowedSymbols[index]
	}

	acc.Password = string(res)
}

func (acc *Account) OutputAccInfo() {
	fmt.Println(acc.Login)
	fmt.Println(acc.Password)
	fmt.Println(acc.Url)
}
