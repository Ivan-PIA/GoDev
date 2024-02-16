package main

import (
	"fmt"
)

func summ(a, b float64) float64 {
	return a + b
}
func add(a, b float64) float64 {
	return a * b
}

func minus(a, b float64) float64 {
	return a - b
}

func delen(a, b float64) float64 {
	return a / b
}

func main() {
	var first float64
	var second float64
	var oper string
	for {
		fmt.Print("Введите первое число: ")
		_, err := fmt.Scan(&first)

		if err != nil {
			fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
		} else {
			break
		}

	}
	//fmt.Print("\n")

	for {
		fmt.Print("Выберите операцию: (+, -, *, /): ")
		fmt.Scanf("%s", &oper)
		if oper == "+" {
			break
		} else if oper == "-" {
			break
		} else if oper == "*" {
			break
		} else if oper == "/" {
			break
		} else {
			fmt.Println("Некорректная операция. Пожалуйста, используйте символы +, -, * или /")
		}
	}

	//fmt.Print("\n")
	for {
		fmt.Print("Введите второе число: ")

		_, err2 := fmt.Scan(&second)

		if err2 != nil {
			fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
		} else if oper == "/" && second == 0 {
			fmt.Println("Ошибка: деление на ноль невозможно.")
		} else {
			break
		}

	}
	//fmt.Print("\n")

	var res float64

	if oper == "+" {
		res = summ(first, second)
	} else if oper == "-" {
		res = minus(first, second)
	} else if oper == "*" {
		res = add(first, second)
	} else if oper == "/" {
		res = delen(first, second)
	} else {
		fmt.Println("dsds")
	}

	fmt.Println("Ответ: ", res)
}
