package main

import (
	"fmt"
	"strconv"
)

func countStringNumber(operation string) int {
	var result int
	var operand int
	for _, val := range operation {
		// fmt.Println(val)
		if val >= 48 && val <= 57 {
			i, _ := strconv.Atoi(string(val))
			if operand == 43 {
				result += i
				operand = 0
			} else if operand == 45 {
				result -= i
				operand = 0
			} else {
				result += i
				operand = 0
			}
			// fmt.Println(num)
		} else if val == 43 || val == 45 {
			operand = int(val)
		}
		// fmt.Println(result)

	}
	// fmt.Println(operand)
	return result
}

func main() {
	operation := "-5 + 2 - 1 - 2"
	asek := countStringNumber(operation)
	fmt.Println(asek)
}
