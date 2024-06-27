package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Альтернативное задание.")
	fmt.Println("Введите строку вида 'a' + 'b', 'a' - 'b', 'a' * b, 'a' / b:")

	string_request := bufio.NewReader(os.Stdin)
	saving_the_string, err := string_request.ReadString('\n')
	saving_the_string = strings.TrimSpace(saving_the_string)

	if err != nil {
		panic("Error")
	}

	operators := []string{" + ", " - ", " * ", " / "}
	parts, operator := check_operators(saving_the_string, operators)

	if operator == "" {
		panic("Ошибка связана с оператором")
	}

	if len(parts) != 2 {
		panic("Ошибка связана с количеством частей")
	}

	switch operator {
	case " + ":
		answer := calc1(saving_the_string)
		fmt.Println(answer)
	case " - ":
		answer := calc2(saving_the_string)
		fmt.Println(answer)
	case " * ":
		answer := calc3(saving_the_string)
		fmt.Println(answer)
	case " / ":
		answer := calc4(saving_the_string)
		fmt.Println(answer)
	default:
		fmt.Println("Неизвестный оператор")
	}
}

// поиск операторов в введенной строке
func check_operators(saving_the_string string, operators []string) ([]string, string) {
	for _, operator := range operators {
		if strings.Contains(saving_the_string, operator) {
			return strings.SplitN(saving_the_string, operator, 2), operator
		}
	}
	return []string{saving_the_string}, ""
}

// конкатенация
func calc1(saving_the_string string) string {
	parts := strings.SplitN(saving_the_string, " + ", 2)
	if len(parts) != 2 {
		panic("Ошибка связана с количеством частей в calc1")
	}
	fs := parts[0]
	ss := parts[1]

	res := fs + " " + ss
	return res
}

// вычитание
func calc2(saving_the_string string) string {
	parts := strings.SplitN(saving_the_string, " - ", 2)
	if len(parts) != 2 {
		panic("Ошибка связана с количеством частей в calc2")
	}

	fs := parts[0]
	ss := parts[1]

	res := strings.Replace(fs, ss, "", -1)
	return res
}

// умножение
func calc3(saving_the_string string) string {
	parts := strings.SplitN(saving_the_string, " * ", 2)
	if len(parts) != 2 {
		panic("Ошибка связана с количеством частей в calc3")
	}

	fs := parts[0]
	ss := parts[1]

	ss_to_num, err := strconv.Atoi(ss)
	if err != nil {
		panic("Ошибка связана с конвертацией")
	}

	var res string
	for i := 0; i < ss_to_num; i++ {
		res += fs
	}
	return res
}

// деление
func calc4(saving_the_string string) string {
	parts := strings.SplitN(saving_the_string, " / ", 2)
	if len(parts) != 2 {
		panic("Ошибка связана с количеством частей в calc4")
	}

	fs := parts[0]
	ss := parts[1]

	v_fs := len(fs)

	c_ss, err := strconv.Atoi(ss)
	if err != nil {
		panic("Ошибка конвертации в calc4")
	}

	if c_ss == 0 {
		panic("Деление на ноль")
	}

	pre_res := v_fs / c_ss
	res := fs[:pre_res]
	return res
}
