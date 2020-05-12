package main

import (
	"fmt"
)

//First task - algorithm Luna
func CheckCardNumber(cardNumber string) bool {
	numbers := ConvertCardString(&cardNumber)

	return CalcLun(numbers)
}

func CalcLun(numbers []int) bool {
	var sum int

	if len(numbers)%2 == 0 {
		CalcSum(0, numbers)
	} else {
		CalcSum(1, numbers)
	}

	for _, digit := range numbers {
		sum += digit
	}

	if sum%10 == 0 {
		return true
	}

	return false
}

func CalcSum(pos int, numbers []int) {
	for numIndex := pos; numIndex < len(numbers); {
		x := &numbers[numIndex]

		if v := 2 * *x; v > 9 {
			*x = v - 9
		} else {
			*x = v
		}

		numIndex += 2
	}
}

func ConvertCardString(cardNumber *string) (numbers []int) {
	for _, oneChar := range []rune(*cardNumber) {
		if oneChar < 0 {
			fmt.Println("Invalid card number!")
			break
		}

		numbers = append(numbers, int(oneChar-'0'))
	}
	return
}

//Second task - Hemming distance
func CalcDistsance(firstString, secondString string) (dist int) {
	if len(firstString) != len(secondString) {
		return 0
	}

	firstStringRunes := []rune(firstString)
	secondStringRunes := []rune(secondString)

	for i, oneChar := range firstStringRunes {
		if secondStringRunes[i] != oneChar {
			dist++
		}
	}

	return
}
