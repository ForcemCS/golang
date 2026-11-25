package main

import "fmt"

func main17() {
	a := [5]int{1, 2, 3, 4, 5}

	b := a[1:4] //b 此时是一个切片， 与a建立了关联关系
	fmt.Printf("%#v\n", b)

	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &a[0])
	fmt.Printf("%p\n", &a[1])

	fmt.Printf("%p", &b[0])
}
