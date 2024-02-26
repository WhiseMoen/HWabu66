package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		task := r.URL.Query().Get("task")

		if task == "" {
			http.Error(w, "Отсутствует параметр 'task'", http.StatusBadRequest)
			return
		}

		result, err := calculate(task)
		if err != nil {
			http.Error(w, fmt.Sprintf("Ошибка: %s", err), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "%.2f", result)
	})

	fmt.Println("HTTP-сервер запущен на порту 8080")
	http.ListenAndServe(":8080", nil)
}

func calculate(input string) (float64, error) {
	parts := strings.Fields(input)
	if len(parts) != 3 {
		return 0, fmt.Errorf("неверный формат запроса")
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
