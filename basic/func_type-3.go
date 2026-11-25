package main

import (
	"fmt"
)

func main36() {

	type AAA func(a, b int, c string, d bool) (int, bool)

	var f1, f2 func(a, b int, c string, d bool) (int, bool)

	f1 = func(a, b int, c string, d bool) (int, bool) {

		return a + b, c == "abc" && d
	}

	n, ok := f1(1, 2, "abc", true)

	fmt.Println(n)
	fmt.Println(ok)

	// ----------

	type ConnectionPool struct {
		Server       []string
		LoadBalancer func(a, b int, c string, d bool) (int, bool)
	}

	cp := ConnectionPool{
		Server: []string{"127.0.0.1:1234", "127.0.0.1:5678"},
		//LoadBalancer: f1, //不可以直接使用f2,会报空指针异常
		LoadBalancer: AAA(f1),
	}
	_ = cp
	_ = f2

}
