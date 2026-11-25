package main

import "fmt"

func main8() {
	score := 50

	if score > 90 {
		fmt.Println("优秀")
	} else if  score > 80 {
		fmt.Println("良好")
	} else if  score > 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}
}