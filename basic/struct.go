package main

import "fmt"

// 声明结构体没有赋值的时候，成员变量都会获得相应的零值
// 字符串的0值是"",不是nil
type User struct {
	Name   string
	Age    int
	Score  float64
	Enable bool
}

//结构体方法
func (u User) Introduce() {
	fmt.Printf("我叫%s, 我今年%d岁\n", u.Name, u.Age)
}

func (u User) Age_add() int {
	u.Age++
	return u.Age
}

func main3() {

	var user1 User
	fmt.Printf("%+v\n", user1)

	// 结构体初始化
	user1 = User{
		Name:   "wu kui",
		Age:    24,
		Score:  88.9,
		Enable: true,
	}

	fmt.Printf("%+v\n", user1)

	// 结构体默认类型是值类型，把一个结构体复制给另一个结构体时，会复制一份数据
	// 修改user2,不会影响user1
	user2 := user1
	fmt.Printf("%+v\n", user2)

	user2.Name = "Anna"
	fmt.Printf("%+v\n", user2)
	fmt.Printf("%+v\n", user1)

	// 结构体指针
	user3 := &user1
	fmt.Printf("%+v\n", user3)
	fmt.Println((*user3).Name)
	fmt.Println(user3.Name)

	user4 := new(User) //向操作系统申请一端内存空间，容纳User这种结构体，返回结构体对应的指针
	fmt.Println(user4)
	fmt.Println("--------------")

	//使用结构体的方法
	user3.Introduce()

	//结构体的值复制，在方法内部的作用域完成了++的操作
	new_age := user3.Age_add()
	fmt.Println(new_age)
	// 但是出了方法内部，就是原始变量的值
	fmt.Printf("%d\n", user3.Age)

	// 匿名结构体
	// 结构体类型的变量
	var student struct {
		Name string
		Age  int
	}

	student.Name = "wukui"
	student.Age = 30

}
