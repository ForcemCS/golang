package main

import (
	"errors"
	"fmt"
)

// 接口是一组行为规范的集合
// 接口名通常以er结尾

type Transporter interface {
	move(src string, dest string) (int, error) // 方法名和返回值列表
	whistle(int) int
}

type Steamer interface {
	// move(string, string) (int, error)
	// whistle(int) int
	Transporter
	displacement() int
}

type Car struct {
	name  string
	price int
}

// Car 实现 move
func (c Car) move(src string, dest string) (int, error) {
	if src == dest {
		return 0, errors.New("起点和终点不能相同")
	}

	fmt.Printf("%s 从 %s 开往 %s\n", c.name, src, dest)

	// 假设返回行驶距离
	return 100, nil
}

func (c Car) whistle(n int) int {
	fmt.Printf("%s 鸣笛 %d 次：滴滴滴\n", c.name, n)

	return n
}

func transport(t Transporter) {
	distance, err := t.move("北京", "上海")
	if err != nil {
		fmt.Println("运输失败：", err)
		return
	}

	fmt.Println("运输距离：", distance)

	t.whistle(2)
}

func main32() {

	car := Car{
		name:  "宝马",
		price: 500000,
	}

	transport(car)

}
