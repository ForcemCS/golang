package main

import "fmt"

type Response struct {
	Body string
}

func handler(w *Response) {
	w.Body = "Hello World"
}

func main31() {
	response := Response{Body: ""}

	handler(&response)

	fmt.Println(response.Body)
}
