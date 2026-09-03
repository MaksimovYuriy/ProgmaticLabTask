package main

import (
	"fmt"
)

const digits = "1234567890"

func putSigns(position int, expression []byte, target int) {
	if position == len(digits) {
		if calculate(expression) == target {
			fmt.Println(string(expression))
		}
		return
	}

	digit := digits[position]
	length := len(expression)

	expression = append(expression, '+', digit)
	putSigns(position+1, expression, target)
	expression = expression[:length]

	expression = append(expression, '-', digit)
	putSigns(position+1, expression, target)
	expression = expression[:length]

	expression = append(expression, digit)
	putSigns(position+1, expression, target)
}

func calculate(expression []byte) int {
	result, number, sign := 0, 0, 1

	for _, symbol := range expression {
		if symbol >= '0' && symbol <= '9' {
			number = number*10 + int(symbol-'0')
			continue
		}

		result += sign * number
		number = 0
		if symbol == '+' {
			sign = 1
		} else {
			sign = -1
		}
	}

	return result + sign*number
}

func main() {
	expression := make([]byte, 1, len(digits)*2-1)
	expression[0] = digits[0]
	putSigns(1, expression, 200)
}
