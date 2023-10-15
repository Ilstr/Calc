package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Мапа из Римских чисел
var romanNumbers = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	"XL": 40, "L": 50, "XC": 90, "C": 100,
}

// Конвертация из арабских в римские числа
func convertToRome(arabic int) string {
	val := arabic
	romanNumerals := []struct {
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

	for _, numeral := range romanNumerals {
		for val >= numeral.Value {
			result.WriteString(numeral.Symbol)
			val -= numeral.Value
		}
	}

	return result.String()
}

// Начало
func main() {
	//Ввод, чтение строки
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	// Проверка на наличие оператора
	if !strings.ContainsAny(input, "+-*/") {
		fmt.Println("Ошибка: строка не является математической операцией.")
		return
	}
	// Проверка системы счисления
	if strings.ContainsAny(input, "IVXLC") && regexp.MustCompile(`[0-9]`).MatchString(input) {
		fmt.Println("Ошибка: используются одновременно разные системы счисления.")
		return
	}
	// Проверка на формат математической операции
	operands := strings.Split(input, " ")
	if len(operands) != 3 {
		fmt.Println("Ошибка: формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		return
	}
	// Проверка первого операнда
	errMsg := "Ошибка: числа должны быть римскими или арабскими."
	a, err := strconv.Atoi(operands[0])
	if err != nil {
		val, ok := romanNumbers[operands[0]]
		if !ok {
			fmt.Println(errMsg)
			return
		}
		a = val
	}
	// Проверка второго операнда
	b, err := strconv.Atoi(operands[2])
	if err != nil {
		val, ok := romanNumbers[operands[2]]
		if !ok {
			fmt.Println(errMsg)
			return
		}
		b = val
	}
	// Проверка на целые числа 1 - 10 (I - X)
	if (a > 10 || a < 1) || (b > 10 || b < 1) {
		fmt.Println("Ошибка: число должно быть целым, от 1 до 10 включительно.")
		return
	}
	// Математическая операция +, -, *, /
	operator := operands[1]
	var result int
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	}
	// Проверка на отрицательные числа в римской системе счисления
	if strings.ContainsAny(input, "IVXLC") && result <= 0 {
		fmt.Println("Ошибка: в римской системе счисления нет отрицательных чисел и нуля.")
		return
	}
	// Выводит результат
	if strings.ContainsAny(input, "IVXLC") {
		romeResult := convertToRome(result)
		fmt.Println(romeResult)
	} else {
		fmt.Println(result)
	}
}
