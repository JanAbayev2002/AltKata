package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// тест проекта на другом пк
var max_length = 40

func main() {
	fmt.Println("Альтернативное задание.")
	for {
		fmt.Println("Введите строку вида 'a' + 'b', 'a' - 'b', 'a' * b, 'a' / b:")

		string_request := bufio.NewReader(os.Stdin)
		saving_the_string, err := string_request.ReadString('\n')
		saving_the_string = strings.TrimSpace(saving_the_string)

		if saving_the_string == "exit" {
			fmt.Println("Have a good day, man (i hate girls)")
			os.Exit(0)
		}

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

		// fmt.Println("Проверка, что есть части для программы?", parts[0], "_", parts[1])

		switch operator {
		case " + ":
			answer := calc1(parts)
			fmt.Println(answer)
		case " - ":
			answer := calc2(parts)
			fmt.Println(answer)
		case " * ":
			answer := calc3(parts)
			fmt.Println(answer)
		case " / ":
			answer := calc4(parts)
			fmt.Println(answer)
		default:
			fmt.Println("Неизвестный оператор")
		}
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
func calc1(parts []string) string {

	fs := parts[0]
	ss := parts[1]

	fs = strings.TrimSpace(fs)
	ss = strings.TrimSpace(ss)

	if len(fs) > 10 || len(ss) > 10 {
		fmt.Printf("Строки длиннеее 10 символов"+"|%s||%s|\n", fs, ss)
		panic("Error")
	}

	res := fs + " " + ss
	return res
}

// вычитание
func calc2(parts []string) string {

	fs := parts[0]
	ss := parts[1]

	fs = strings.TrimSpace(fs)
	ss = strings.TrimSpace(ss)

	if len(fs) > 10 || len(ss) > 10 {
		fmt.Printf("Строки длиннеее 10 символов"+"|%s||%s|\n", fs, ss)
		panic("Error")
	}

	res := strings.Replace(fs, ss, "", -1)
	return res
}

// умножение
func calc3(parts []string) string {

	fs := parts[0]
	ss := parts[1]

	fs = strings.TrimSpace(fs)
	ss = strings.TrimSpace(ss)

	if len(fs) > 10 || len(ss) > 10 {
		fmt.Printf("Строки длиннеее 10 символов"+"|%s||%s|\n", fs, ss)
		panic("Error")
	}

	ss_to_num, err := strconv.Atoi(ss)
	if err != nil {
		panic("Ошибка связана с конвертацией")
	}

	if ss_to_num < 1 || ss_to_num > 10 {
		panic("Ошибка. Не соответствует диапазону")
	}

	var res string
	for i := 0; i < ss_to_num; i++ {
		res += fs
	}

	if len(res) > max_length {
		res = res[:max_length] + "..."
	}
	return res
}

// деление
func calc4(parts []string) string {

	fs := parts[0]
	ss := parts[1]

	fs = strings.TrimSpace(fs)
	ss = strings.TrimSpace(ss)

	if len(fs) > 10 || len(ss) > 10 {
		fmt.Printf("Строки длиннеее 10 символов"+"|%s||%s|\n", fs, ss)
		panic("Error")
	}

	v_fs := len(fs)

	c_ss, err := strconv.Atoi(ss)
	if err != nil {
		panic("Ошибка конвертации в calc4")
	}

	if c_ss == 0 {
		panic("Деление на ноль")
	}

	if c_ss < 1 || c_ss > 10 {
		panic("Ошибка. Не соответствует диапазону")
	}

	pre_res := v_fs / c_ss
	res := fs[:pre_res]
	return res
}
