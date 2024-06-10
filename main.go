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
	"X":    10, "XI": 11, "XII": 12, "XIII": 13, "XIV": 14, "XV": 15, "XVI": 16, "XVII": 17, "XVIII": 18, "XIX": 19, "XX": 20, "XXI": 21, "XXII": 22, "XXIV": 24,
	"XXV": 25, "XXVI": 26, "XXIX": 29, "XXX": 30, "XXXI": 31, "XXXV": 35, "XL": 40, "XLI": 41, "XLII": 42, "XLIV": 44, "XLV": 45, "XLVI": 46, "XLIX": 49, "L": 50,
	"LX": 60, "LXX": 70, "LXXX": 80, "XC": 90, "C": 100, "CD": 400, "D": 500, "DC": 600, "CM": 900, "M": 1000,
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
	10: "X", 11: "XI", 12: "XII", 13: "XIII", 14: "XIV", 15: "XV", 16: "XVI", 17: "XVII", 18: "XVIII", 19: "XIX", 20: "XX", 21: "XXI", 22: "XXII", 23: "XXIII",
	24: "XXIV", 25: "XXV", 26: "XXVI", 27: "XXVII", 28: "XXVIII", 29: "XXIX", 30: "XXX", 32: "XXXII", 35: "XXXV", 36: "XXXVI", 40: "XL", 42: "XLII", 45: "XLV",
	48: "XLVIII", 49: "XLIX", 50: "L", 54: "LIV", 56: "LVI", 60: "LX", 63: "LXIII", 64: "LXIV", 70: "LXX", 72: "LXXII", 80: "LXXX", 81: "LXXXI", 90: "XC", 100: "C",
	400: "CD", 500: "D", 600: "DC", 900: "CM", 1000: "M",
}

func rimdig(s string) bool {
	for _, c := range s {
		if _, ok := rimdigits[string(c)]; !ok {
			return false
			/*if rimdigits[s] == 0 {
				return false
			}*/
		}
	}
	return true

}

var arabdigits = map[string]int{
	"0":  0,
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
	/*if arabdigits[s] == 0 {
	return false*/
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
	var trash string
	var result int

	fmt.Println("Введите выражение: ")
	fmt.Scanln(&operand1, &operator, &operand2, &trash)

	for rimdig(operand1) {
		if rimdig(operand1) && rimdig(operand2) {
			operand1 = strings.TrimSpace(operand1)
			operand2 = strings.TrimSpace(operand2)

			if rimdigits[operand1] <= 10 && rimdigits[operand2] <= 10 && /*len(operand1) <= 4 && len(operand2) <= 4 &&*/ len(trash) < 1 && rimdigits[operand1] != 0 && rimdigits[operand2] != 0 {
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
					if result >= 1 {
						fmt.Println(stigidmir[result])
					} else {
						panic("Результатом деления является только целое число")
					}

				default:
					panic("Строка не является математической операцией!!!")
				}
			} else {
				panic("Калькулятор принимает на ввод только 2 числа от I до X включительно")
			}
		} else {
			panic("Используются 2 разные системы счисления")
		}
		break
	}

	for arabdig(operand2) {
		if arabdig(operand1) && arabdig(operand2) {

			operand1 = strings.TrimSpace(operand1)
			operator1, _ := strconv.Atoi(operand1)
			operand2 = strings.TrimSpace(operand2)
			operator2, _ := strconv.Atoi(operand2)
			if operator1 <= 10 && operator2 <= 10 /*&& operator1 != 0 && operator2 != 0*/ && len(trash) < 1 {

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
					panic("Строка не является математической операцией!!!")
				}
			} else {
				panic("Калькулятор принимает на ввод только 2 числа от 1 до 10 включительно")
			}
		} else {
			panic("Используются 2 разные системы счисления")
		}
		break
	}

}
