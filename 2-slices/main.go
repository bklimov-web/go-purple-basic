package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var numbersStr string

	operation := getOperation()

	fmt.Print("Введите числа через запятую: ")
	fmt.Scan(&numbersStr)

	numbers, err := getNumbers(numbersStr)

	if err != nil {
		fmt.Print(err)
	}

	if operation == "SUM" {
		fmt.Print(calcSum(numbers))
	} else if operation == "AVG" {
		fmt.Print(calcAvg(numbers))
	} else {
		fmt.Print(calcMed(numbers))
	}
}

func getNumbers(numbersStr string) ([]int, error) {
	strNumbers := strings.Split(numbersStr, ",")
	numbers := make([]int, 0, len(strNumbers))

	for _, value := range strNumbers {
		intValue, err := strconv.Atoi(strings.TrimSpace(value))

		if err != nil {
			return []int{}, errors.New("PARSE_INT_ERROR")
		}

		numbers = append(numbers, intValue)
	}

	return numbers, nil
}

func getOperation() string {
	var operation string
	fmt.Print("Введите операцию (SUM, AVG, MED): ")
	fmt.Scan(&operation)

	for {
		if operation != "SUM" && operation != "AVG" && operation != "MED" {
			fmt.Println("Не подходящая операция")
			fmt.Print("Введите операцию (SUM, AVG, MED): ")
			fmt.Scan(&operation)
			continue
		}
		break
	}

	return operation
}

func calcSum(numbers []int) int {
	var sum int

	for _, value := range numbers {
		sum += int(value)
	}

	return sum
}

func calcAvg(numbers []int) int {
	avg := calcSum(numbers) / len(numbers)

	return avg
}

func calcMed(numbers []int) int {
	length := len(numbers)
	sort.Ints(numbers)

	if length%2 == 0 {
		return (numbers[length/2-1] + numbers[length]) / 2
	} else {
		return numbers[length/2]
	}
}
