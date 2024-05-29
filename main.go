package main

import (
	"fmt"
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
}

var stigidmir = map[int]string{
	1:  "I",
	2:  "II",
	3:  "III",
	4:  "IV",
	5:  "V",
	6:  "VI",
	7:  "VII",
	8:  "VIII",
	9:  "IX",
	10: "X", 40: "XL", 50: "L", 90: "XC", 100: "C", 400: "CD", 500: "D", 600: "DC", 900: "CM", 1000: "M",
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

	fmt.Println("Введите выражение: ")
	fmt.Scanln(&operand1, &operator, &operand2)

	for rimdig(operand1) {
		if rimdig(operand2) {
			operand1 = strings.TrimSpace(operand1)
			//operator1, _ := strconv.Atoi(operand1)

			operand2 = strings.TrimSpace(operand2)
			//operator2, _ := strconv.Atoi(operand2)
			switch operator {
			case "+":
				result = rimdigits[operand1] + rimdigits[operand2]
				fmt.Println(stigidmir[result])

			case "-":
				result = rimdigits[operand1] - rimdigits[operand2]
				if result > 0 {
					fmt.Println(stigidmir[result])
				} else {
					panic("В римской системе счисления нет отрицательных чисел")
				}

			case "*":
				result = rimdigits[operand1] * rimdigits[operand2]
				fmt.Println(stigidmir[result])

			case "/":
				result = rimdigits[operand1] / rimdigits[operand2]
				if result > 1 {
					fmt.Println(stigidmir[result])
				} else {
					panic("Результатом деления является только целое число")
				}

			default:
				fmt.Println("Строка не является математической операцией!!!")

			}
		} else {
			panic("Используются 2 разные системы счисления")
		}
		break
	}

	for arabdig(operand1) {
		if arabdig(operand2) {
			operand1 = strings.TrimSpace(operand1)
			//operator1, _ := strconv.Atoi(operand1)

			operand2 = strings.TrimSpace(operand2)
			//operator2, _ := strconv.Atoi(operand2)
			switch operator {
			case "+":
				result = arabdigits[operand1] + arabdigits[operand2]
				fmt.Println(result)

			case "-":
				result = arabdigits[operand1] - arabdigits[operand2]
				fmt.Println(result)

			case "*":
				result = arabdigits[operand1] * arabdigits[operand2]
				fmt.Println(result)

			case "/":
				result = arabdigits[operand1] / arabdigits[operand2]
				fmt.Println(result)

			default:
				fmt.Println("Строка не является математической операцией!!!")
			}
		} else {
			panic("Используются 2 разные системы счисления")
		}
		break

	}
}
