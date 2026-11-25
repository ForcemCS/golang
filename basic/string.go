package main

import (
	"fmt"
	"strings"
)

func main22() {
	s1 := `hello "golang"`
	fmt.Println(s1)

	// byte相当于unit8
	// rune相当于nit32

	var a []byte = []byte{65, 101, 108, 108, 111}
	fmt.Printf("%c\n", a[0])

	s2 := "My name is  吴魁"
	arr := []byte(s2)
	fmt.Printf("%c\n", arr[0])

	brr := []rune(s2)
	fmt.Printf("%c\n", brr[0])

	for _, ele := range s2 {
		fmt.Printf("%c\n", ele)
	}

	//字符串的拼接

	uid := "10001"
	key1 := "user:" + uid

	key2 := fmt.Sprintf("%s%s", "user:", uid)
	fmt.Println(key1)
	fmt.Println(key2)

	names := []string{
		"alice",
		"bob",
		"tom",
	}

	result := strings.Join(names, ":")
	fmt.Println(result)

}
