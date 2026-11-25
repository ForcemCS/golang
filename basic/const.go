package main

import "fmt"

func main2() {

	// const (
	// 	a = iota
	// 	b
	// 	c = 7
	// 	d = 2
	// 	e       // 复制d
	// 	f       // 复制e
	// )

	// fmt.Println(f)

	const (
		NOT_USE = 1 << (10 * iota)
		KB      // 复制上一行的表达式
		MB
		GB
	)
	fmt.Println(KB)

}
