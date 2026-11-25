package main

import (
	"fmt"
	"strconv" //字符串和其他类型之间的转换，不是字符
)

func main23() {
	s1 := "9"

	n1, err1 := strconv.Atoi(s1)
	if err1 != nil {
		fmt.Println("转换失败: ", err1)
	}

	n2, err2 := strconv.ParseInt(s1, 10, 64)
	if err2 != nil {
		return
	}

	fmt.Println(n1)
	fmt.Printf("%T\n", n1)
	fmt.Println(n2)
	fmt.Printf("%T\n", n2)

	n3 := 123

	s2 := strconv.Itoa(n3)
	fmt.Println(s2)
	fmt.Printf("%T\n", s2)
}
