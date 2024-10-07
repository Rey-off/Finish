package main

import (
	"fmt"
	"strconv"
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

var romanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

type Number struct {
	Value int
	Type  string
}

func main() {
	for {
		fmt.Print("Введите выражение: ")
		var input string
		fmt.Scanln(&input)

		input = strings.ReplaceAll(input, " ", "")

		if len(input) < 3 {
			fmt.Println("Неправильный ввод. Пожалуйста, введите правильное выражение, например '2+3' или 'II+III'.")
			continue
		}

		var num1, num2 string
		var op string
		for i, c := range input {
			if c == '+' || c == '-' || c == '*' || c == '/' {
				num1 = input[:i]
				op = string(c)
				num2 = input[i+1:]
				break
			}
		}

		if num1 == "" || op == "" || num2 == "" {
			fmt.Println("Неправильный ввод. Пожалуйста, введите правильное выражение, например '2+3' или 'II+III'.")
			continue
		}

		num1Num, err := parseNumber(num1)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if op != "+" && op != "-" && op != "*" && op != "/" {
			fmt.Println("Паника:Неправильный оператор")
			continue
		}

		num2Num, err := parseNumber(num2)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if num1Num.Type != num2Num.Type {
			fmt.Println("Паника:Смешанные типы чисел")
			continue
		}

		var result int
		switch op {
		case "+":
			result = num1Num.Value + num2Num.Value
		case "-":
			result = num1Num.Value - num2Num.Value
		case "*":
			result = num1Num.Value * num2Num.Value
		case "/":
			if num2Num.Value == 0 {
				fmt.Println("Паника:Деление на ноль")
				continue
			}
			result = num1Num.Value / num2Num.Value
		}

		if result < 0 && num1Num.Type == "roman" {
			fmt.Println("Паника: Римские цифры не могут быть отрицательными")
			continue
		}

		if num1Num.Type == "arabic" {
			fmt.Println(result)
		} else {
			romanResult, err := toRoman(result)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(romanResult)
		}
	}
}

func parseNumber(s string) (Number, error) {
	if s == "I" || s == "II" || s == "III" || s == "IV" || s == "V" || s == "VI" || s == "VII" || s == "VIII" || s == "IX" || s == "X" {
		return Number{toArabic(s), "roman"}, nil
	} else if i, err := strconv.Atoi(s); err == nil && i >= 1 && i <= 10 {
		return Number{i, "arabic"}, nil
	} else {
		return Number{}, fmt.Errorf("Неправильное число")
	}
}

func toArabic(s string) int {
	switch s {
	case "I":
		return 1
	case "II":
		return 2
	case "III":
		return 3
	case "IV":
		return 4
	case "V":
		return 5
	case "VI":
		return 6
	case "VII":
		return 7
	case "VIII":
		return 8
	case "IX":
		return 9
	case "X":
		return 10
	default:
		return 0
	}
}

func toRoman(i int) (string, error) {
	if i < 1 {
		return "", fmt.Errorf("Неправильный результат")
	} else if i == 0 {
		return "0", nil
	}

	var result string
	for _, rn := range romanNumerals {
		for i >= rn.Value {
			result += rn.Symbol
			i -= rn.Value
		}
	}
	return result, nil
}
