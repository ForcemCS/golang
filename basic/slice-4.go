package main

import "fmt"

func main20() {

	var arr [3]int

	brr := arr[1:2]

	fmt.Println(len(arr))
	fmt.Println(len(brr), cap(brr))

	brr = append(brr, 9)
	fmt.Println(arr, brr)

	brr = append(brr, 9)
	fmt.Println(cap(brr))

}