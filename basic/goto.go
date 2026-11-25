package main

import "fmt"

func baisc_goto() {
	var i int = 4
MY_LABEL:
	i += 3
	i *= 2
	fmt.Println(i)

	if i > 200 {
		return
	}

	goto MY_LABEL

}

func main11() {
	baisc_goto()

}
