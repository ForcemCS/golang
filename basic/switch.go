package main

import "fmt"

// 此时是短路执行，满足一个条件就终止了
func basic_switch() {
	color := "yellow"

	// switch color {
	// case "green":
	// 	fmt.Println("gooo")
	// case "yellow", "red":
	// 	fmt.Println("ok")
	// default:
	// 	fmt.Println("STOP")
	// }

	switch {
	case color == "yellow":
		fmt.Println("OK!")
	default:
		fmt.Println("STOP")
	}

}

func main14() {
	basic_switch()

}
