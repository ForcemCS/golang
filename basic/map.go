package main

import "fmt"

// key 在作为索引期间，其参与 Hash 和相等比较的内容不能随便变化。
// map 的 key 必须是 comparable（可比较类型）。

// map[int]string
// map[int64]string
// map[string]int
// map[bool]string

type Positon struct {
	X int
	Y int
}

func main24() {

	// 数组可以当作key
	m1 := map[[3]int]string{}
	m1[[3]int{1, 2, 3}] = "hello"
	fmt.Printf("%v\n", m1)

	//结构体也可以作为key

	m2 := map[Positon]string{}
	m2[Positon{X: 10, Y: 20}] = "boss"
	fmt.Println(m2[Positon{10, 20}])

	//----------
	m3 := make(map[int]string, 5) //初始化建议制定容量
	m3[1] = "hello"
	fmt.Println(m3)

	m4 := map[string]int{"数学": 100, "物理": 80}
	fmt.Println(m4["数学"])
	m4["语文"] = 60
	m4["数学"] = 150
	fmt.Println(m4)
	fmt.Println(m4["化学"]) // 不存在的话，返回零值

	delete(m4, "数学")

	v, ok := m4["数学"]
	fmt.Println(ok, v) //key不存在

	//------需要注意的是：即使没有扩容，返回的键值对的顺序，也是无序的。但是是一个环
	for keys, value := range m4 {
		fmt.Println(keys, value)
	}

}
