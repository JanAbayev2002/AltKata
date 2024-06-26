package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Альтернативное задание.")
	fmt.Println("Введите строку вида 'a' + 'b', 'a' - 'b', 'a' * b, 'a' / b:")

	input := bufio.NewReader(os.Stdin)
	saved, err := input.ReadString('\n')

	if err != nil {
		panic(err)
	}

	saved = strings.TrimSpace(saved)

	if saved == "exit" {
		fmt.Println("Вы покинули проект")
		os.Exit(0)
	}

	parts, operator, err := split_input(saved)
	if err != nil {
		panic(err)
	}

	fmt.Println(parts[0])
	fmt.Println(operator)
	fmt.Println(parts[1])
}

func split_input(input string) ([2]string, string, error) {

	operators := []string{" + ", " - ", " * ", " / "}

	for _, operator := range operators {
		if strings.Contains(input, operator) {
			parts := strings.SplitN(input, operator, 2)

			//fmt.Println(parts[0], operator, parts[1])

			for _, oper := range operators {
				if strings.Contains(parts[1], oper) {
					panic("Неправильное выражение")
				}
			}

			return [2]string{strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])}, operator, nil
		}
	}

	panic("Оператор не найден")
}

//  func calc(parts []string, operator string) (string, error) {
// 	array := [3]string {parts[0], operator, parts[1]}
//  }
