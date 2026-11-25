package main

import (
	"fmt"
)

type EmptyInterface interface{}

func Sum1(args ...EmptyInterface) int {
	rect := 0

	for _, ele := range args {
		switch v := ele.(type) {
		case int:
			rect += v
		case float32:
			rect += int(v)
		default:
			fmt.Printf("不支持的数据类型%T 值为%v\n", ele, ele)
		}
	}

	return rect
}

func main37() {
	rect := Sum1(1, float32(3.14), "aaa")
	fmt.Println(rect)

	// ----------------

	//类型断言

	var a any

	a = "ccc"

	if v, ok := a.(int); ok {
		fmt.Printf("%d\n", v)
	}
}
