package main

import (
	"errors"
	"fmt"
)

func main() {
	amount, initialCurrency, targetCurrency, err := getUserInput()

	if err != nil {
		fmt.Printf("Ошибка: %v", err)
		return
	}

	result := calculateTarget(amount, initialCurrency, targetCurrency)
	fmt.Printf("%v %v = %.2f %v", amount, initialCurrency, result, targetCurrency)
}

func getUserInput() (float64, string, string, error) {
	var amount float64
	var initialCurrency string
	var targetCurrency string

	fmt.Printf("Введите исходную валюту (USD, EUR, RUB): ")
	fmt.Scan(&initialCurrency)

	if initialCurrency != "USD" && initialCurrency != "EUR" && initialCurrency != "RUB" {
		return 0, "", "", errors.New("CURRENCY_ERROR")
	}

	fmt.Printf("Введите число: ")
	fmt.Scan(&amount)

	if amount <= 0 {
		return 0, "", "", errors.New("MISSING_AMOUNT")
	}

	fmt.Printf("Введите целевую валюту %v: ", getCurrencyHint(initialCurrency))
	fmt.Scan(&targetCurrency)

	if initialCurrency != "USD" && initialCurrency != "EUR" && initialCurrency != "RUB" || targetCurrency == initialCurrency {
		return 0, "", "", errors.New("CURRENCY_ERROR")
	}

	return amount, initialCurrency, targetCurrency, nil
}

func getCurrencyHint(initial string) string {
	if initial == "USD" {
		return "(EUR, RUB)"
	}

	if initial == "EUR" {
		return "(USD, RUB)"
	}

	return "(EUR, USD)"
}

func calculateTarget(amount float64, initialCurrency string, targetCurrency string) float64 {
	const USD_EUR = 0.88
	const USD_RUB = 81.25

	if initialCurrency == "USD" {
		if (targetCurrency == "EUR") {
			return amount * USD_EUR
		}
		return amount * USD_RUB
	}

	if initialCurrency == "EUR" {
		if (targetCurrency == "USD") {
			return amount / USD_EUR
		}
		return amount / USD_EUR * USD_RUB
	}

	if (targetCurrency == "USD") {
		return amount / USD_RUB
	}
	return amount / USD_RUB * USD_EUR
}