package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Простой консольный калькулятор. Введите 'exit' для завершения.")

	for {
		fmt.Print("Введите выражение в формате ЧислоОперацияЧисло: ")
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			break
		}

		result, err := calculate(input)
		if err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			fmt.Println("Ответ:", result)
		}
	}
}

func calculate(input string) (float64, error) {
	// Разбиваем введенную строку на число, операцию и второе число
	parts := strings.Fields(input)
	if len(parts) != 3 {
		return 0, fmt.Errorf("неверный формат ввода")
	}

	num1, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, fmt.Errorf("ошибка при чтении первого числа")
	}

	operator := parts[1]

	num2, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return 0, fmt.Errorf("ошибка при чтении второго числа")
	}

	// Выполняем операцию
	var result float64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("деление на ноль запрещено")
		}
		result = num1 / num2
	default:
		return 0, fmt.Errorf("неподдерживаемая операция: %s", operator)
	}

	return result, nil
}
