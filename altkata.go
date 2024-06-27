package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	answer1 := calc1()
	fmt.Println(answer1)

	answer2 := calc2()
	fmt.Println(answer2)

	answer3 := calc3()
	fmt.Println(answer3)

	answer4 := calc4()
	fmt.Println(answer4)

}

// конкатенация
func calc1() string {
	str1 := "Hello, my name is + Yan"
	parts := strings.SplitN(str1, " + ", 2)
	if len(parts) != 2 {
		panic("Error")
	}
	fs := parts[0]
	ss := parts[1]

	res := fs + ss
	return (res)
}

// вычитание
func calc2() string {
	str1 := "Yan and Mila - Mir"
	parts := strings.SplitN(str1, " - ", 2)
	if len(parts) != 2 {
		panic("Error")
	}

	fs := parts[0]
	ss := parts[1]

	res := strings.Replace(fs, ss, "", -1)
	return res
}

// умножение
func calc3() string {
	str1 := "Hello, mama * 4"
	parts := strings.SplitN(str1, " * ", 2)
	if len(parts) != 2 {
		panic("Error")
	}

	fs := parts[0]
	ss := parts[1]

	ss_to_num, err := strconv.Atoi(ss)
	if err != nil {
		panic("Error")
	}

	var res string
	for i := 0; i < ss_to_num; i++ {
		res += fs
	}
	return res
}

// деление
func calc4() string {
	str1 := "Never give up / 4"
	parts := strings.SplitN(str1, " / ", 2)

	if len(parts) != 2 {
		panic("Error")
	}

	fs := parts[0]
	ss := parts[1]

	v_fs := len(fs)

	c_ss, err := strconv.Atoi(ss)

	if err != nil {
		panic("Error")
	}

	if c_ss == 0 {
		panic("Error")
	}

	pre_res := v_fs / c_ss
	res := str1[:pre_res]
	return res
}
