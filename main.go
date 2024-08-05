package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	operand1, operator, operand2, err := parseInput(input)
	if err != nil {
		fmt.Print(err)
		return
	}
	numbericType, err := checkType(operand1, operand2)
	if err != nil {
		fmt.Print(err)
		return
	}
	switch numbericType {
	case "roman":
		{
			res, err := calculateRoman(operand1, operator, operand2)
			if err != nil {
				fmt.Print(err)
				return
			}
			fmt.Print(res)
			return
		}
	case "arabic":
		{
			res, err := calculateArabic(operand1, operator, operand2)
			if err != nil {
				fmt.Print(err)
				return
			}
			fmt.Print(res)
			return
		}
	}
}
func parseInput(input string) (operand1 string, operator string, operand2 string, err error) {
	x := strings.TrimSpace(input)
	parts := strings.Split(x, " ")
	if len(parts) == 1 {
		err = errors.New("Выдача паники, так как строка не является математической операцией.")
		return
	}
	if len(parts) != 3 {
		err = errors.New("Выдача паники, так как формат математической" +
			" операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		return
	}

	operand1 = parts[0]
	operator = parts[1]
	operand2 = parts[2]

	validOperators := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}

	if !validOperators[operator] {
		err = errors.New("Выдача паники, так как формат математической операции не удовлетворяет" +
			" заданию — два операнда и один оператор (+, -, /, *).")
		return
	}
	return
}
func calculate(x int, s string, y int) string {
	var res int
	switch s {
	case "+":
		res = x + y
	case "-":
		res = x - y
	case "*":
		res = x * y
	case "/":
		res = x / y
	}
	return strconv.Itoa(res)
}
func calculateArabic(operand1 string, operator string, operand2 string) (string, error) {
	x, _ := strconv.Atoi(operand1)
	y, _ := strconv.Atoi(operand2)
	if x < 1 || x > 10 || y < 1 || y > 10 {
		return "", errors.New("Выдача паники, операнд не [1,10].")
	}
	res := calculate(x, operator, y)
	return res, nil
}
func convertRomanToArabic(roman string) (string, error) {
	romanToArabic := map[string]string{
		"I":    "1",
		"II":   "2",
		"III":  "3",
		"IV":   "4",
		"V":    "5",
		"VI":   "6",
		"VII":  "7",
		"VIII": "8",
		"IX":   "9",
		"X":    "10",
	}

	if value, exists := romanToArabic[roman]; exists {
		return value, nil
	} else {
		return "", errors.New("Выдача паники, операнд не [1,10].")
	}
}
func convertArabicToRoman(arabic string) (string, error) {
	arabicNumber, _ := strconv.Atoi(arabic)
	vals := []struct {
		Value  int
		Symbol string
	}{
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

	var result strings.Builder
	for _, val := range vals {
		for arabicNumber >= val.Value {
			result.WriteString(val.Symbol)
			arabicNumber -= val.Value
		}
	}
	return result.String(), nil
}
func checkType(operand1 string, operand2 string) (string, error) {
	_, errRoman1 := convertRomanToArabic(operand1)
	_, errRoman2 := convertRomanToArabic(operand2)
	// проверка оба римские
	if errRoman1 == nil && errRoman2 == nil {
		return "roman", nil
	}
	_, errArabic1 := strconv.Atoi(operand1)
	_, errArabic2 := strconv.Atoi(operand2)
	// проверка оба арабские
	if errArabic1 == nil && errArabic2 == nil {
		return "arabic", nil
	}
	// Проверка на использование нецелого числа, например (1.5+int)
	// *если для этого нужна отдельная паника в тз не указано
	_, err := strconv.ParseFloat(operand1, 64)
	if err == nil {
		_, err = strconv.Atoi(operand2)
		if err == nil {
			return "", errors.New("Выдача паники, используется нецелое число")
		}
	}
	_, err = strconv.ParseFloat(operand2, 64)
	if err == nil {
		_, err = strconv.Atoi(operand1)
		if err == nil {
			return "", errors.New("Выдача паники, используется нецелое число")
		}
	}
	if (errRoman1 == nil && errArabic2 == nil) || (errRoman2 == nil && errArabic1 == nil) {
		return "", errors.New("Выдача паники, так как используются одновременно разные системы счисления.")
	}
	if errRoman1 != nil {
		return "", errRoman1
	} else {
		return "", errRoman2
	}
}

func calculateRoman(a string, s string, c string) (string, error) {
	x, _ := convertRomanToArabic(a)
	y, _ := convertRomanToArabic(c)
	res, _ := calculateArabic(x, s, y)
	resInt, _ := strconv.Atoi(res)
	if resInt < 1 {
		return "", errors.New("Римские цифры не могут быть меньше 1")
	}
	res, _ = convertArabicToRoman(res)
	return res, nil
}
