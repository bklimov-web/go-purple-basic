package main

import (
	"fmt"
)

func main() {
	const USD_EUR = 0.88
	const USD_RUB = 81.25
	amount, initialCurrency, targetCurrency := getUserInput()

	const EUR_RUB = USD_RUB / USD_EUR
	fmt.Print(EUR_RUB)
}

func getUserInput() (float64, string, string) {
	var amount float64
	var initialCurrency string
	var targetCurrency string

	fmt.Printf("Введите число: ")
	fmt.Scan(&amount)

	fmt.Printf("Введите исходную валюту: ")
	fmt.Scan(&initialCurrency)

	fmt.Printf("Введите целевую валюту: ")
	fmt.Scan(&targetCurrency)

	return amount, initialCurrency, targetCurrency
}

func calculateTarget(amount float64, initialCurrency string, targetCurrency string) {}