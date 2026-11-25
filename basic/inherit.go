package main

import "fmt"

func mai4() {

	type User struct {
		Name string
		Age  int
	}

	type Video struct {
		Length int
		Name   string
		// Auther User
		User
	}

	u := User{Name: "吴魁", Age: 18}
	video := Video{Length: 129, Name: "GO教程", User: u}
	fmt.Println(video.Age)
	fmt.Println(video.User.Name)

}
