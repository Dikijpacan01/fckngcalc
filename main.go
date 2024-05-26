package main

import (
	"fmt"
	"strconv"
	"strings"
)

var rimdigits = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
	"XL":   40, "L": 50, "XC": 90, "C": 100, "CD": 400, "D": 500, "DC": 600, "CM": 900, "M": 1000,
}

func rimdig(s string) bool {
	for _, c := range s {
		if _, ok := rimdigits[string(c)]; !ok {
			return false
		}
	}
	return true

}

var arabdigits = map[string]int{
	"1":  1,
	"2":  2,
	"3":  3,
	"4":  4,
	"5":  5,
	"6":  6,
	"7":  7,
	"8":  8,
	"9":  9,
	"10": 10,
}

func arabdig(s string) bool {
	for _, c := range s {
		if _, ok := arabdigits[string(c)]; !ok {
			return false
		}
	}
	return true
}

func main() {

	var operator string
	var operand1 string
	var operand2 string
	var result int
	var result1 int

	fmt.Println("Введите выражение: ")
	fmt.Scan(&operand1, &operator, &operand2)

	for rimdig(operand1) {
		if rimdig(operand2) {
			operand1 = strings.TrimSpace(operand1)
			//operator1, _ := strconv.Atoi(operand1)

			operand2 = strings.TrimSpace(operand2)
			//operator2, _ := strconv.Atoi(operand2)
			switch operator {
			case "+":
				result = rimdigits[operand1] + rimdigits[operand2]

			case "-":
				result = rimdigits[operand1] - rimdigits[operand2]

			case "*":
				result = rimdigits[operand1] + rimdigits[operand2]

			case "/":
				result = rimdigits[operand1] / rimdigits[operand2]
			}
		}

		fmt.Println(result)

		break
	}

	for arabdig(operand1) {
		if arabdig(operand2) {
			operand1 = strings.TrimSpace(operand1)
			operator1, _ := strconv.Atoi(operand1)

			operand2 = strings.TrimSpace(operand2)
			operator2, _ := strconv.Atoi(operand2)
			switch operator {
			case "+":
				result1 = operator1 + operator2

			case "-":
				result1 = operator1 - operator2

			case "*":
				result1 = operator1 * operator2

			case "/":
				result1 = operator1 / operator2
			}
			fmt.Println(result1)
		}
		break

	}

	/*operand1 = strings.TrimSpace(operand1)
	operator1, _ := strconv.Atoi(operand1)

	operand2 = strings.TrimSpace(operand2)
	operator2, _ := strconv.Atoi(operand2)*/

	/*operand1, _ := strconv.Atoi(strings.TrimSpace(operand1Str))

	operand2Str, _ := reader.Readstring('n')*/
	//name, _ := strconv.Atoi(strings.TrimSpace(nameStr))
	/*var operand1 int
	var operand2 int
	var result int*/

	/*fmt.Println("Введите свое выражение: ")
	  fmt.Scan(&operand1, &operator, &operand2)

	  switch operator {
	  case "+":
	  	result = operand1 + operand2

	  case "-":
	  	result = operand1 - operand2

	  case "*":
	  	result = operand1 * operand2

	  case "/":
	  	result = operand1 / operand2
	  }
	  fmt.Println(result)*/
}

//fmt.Println(rimdigits[name] + rimdigits[name1])
//fmt.Println(operator + operator2)
