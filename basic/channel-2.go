package main

import "fmt"

func main30() {

	var aa int

	ch1 := make(chan struct{})

	go func() {

		aa = 100

		// 告诉 main：
		// “我已经把 x 改完了”
		ch1 <- struct{}{}

	}()

	// main 在这里等通知
	<-ch1

	// 收到通知以后，才继续
	fmt.Println(aa)

}
