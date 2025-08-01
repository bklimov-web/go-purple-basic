package main

import (
	"fmt"
)

func main() {
	initialCurrency := getInitialCurrency()
	amount := getAmount()
	targetCurrency := getTargetCurrency(initialCurrency)

	result := calculateTarget(amount, initialCurrency, targetCurrency)
	fmt.Printf("%v %v = %.2f %v", amount, initialCurrency, result, targetCurrency)
}

func getInitialCurrency() string {
	var initialCurrency string
	fmt.Print("Введите исходную валюту (USD, EUR, RUB): ")
	fmt.Scan(&initialCurrency)

	for {
		if initialCurrency != "USD" && initialCurrency != "EUR" && initialCurrency != "RUB" {
			fmt.Println("Не подходящая валюта, попробуйте снова")
			fmt.Print("Введите исходную валюту (USD, EUR, RUB): ")
			fmt.Scan(&initialCurrency)
			continue
		}
		break
	}

	return initialCurrency
}

func getTargetCurrency(initialCurrency string) string {
	var targetCurrency string

	fmt.Printf("Введите исходную валюту %v: ", getCurrencyHint(initialCurrency))
	fmt.Scan(&targetCurrency)

	for {
		if (targetCurrency != "USD" && targetCurrency != "EUR" && targetCurrency != "RUB") || targetCurrency == initialCurrency {
			fmt.Println("Не подходящая валюта, попробуйте снова")
			fmt.Printf("Введите исходную валюту %v: ", getCurrencyHint(initialCurrency))
			fmt.Scan(&targetCurrency)
			continue
		}
		break
	}

	return targetCurrency
}

func getAmount() float64 {
	var amount float64
	fmt.Print("Введите сумму: ")
	fmt.Scan(&amount)

	for {
		if amount <= 0 {
			fmt.Println("Сумма неверная, попробуйте снова")
			fmt.Print("Введите сумму: ")
			fmt.Scan(&amount)
			continue
		}
		break
	}

	return amount
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
		if targetCurrency == "EUR" {
			return amount * USD_EUR
		}
		return amount * USD_RUB
	}

	if initialCurrency == "EUR" {
		if targetCurrency == "USD" {
			return amount / USD_EUR
		}
		return amount / USD_EUR * USD_RUB
	}

	if targetCurrency == "USD" {
		return amount / USD_RUB
	}
	return amount / USD_RUB * USD_EUR
}
