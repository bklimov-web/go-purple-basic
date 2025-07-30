package main

import (
	"fmt"
)

func main() {
	const USD_EUR = 0.88
	const USD_RUB = 81.25

	const EUR_RUB = USD_RUB / USD_EUR
	fmt.Print(EUR_RUB)
}
