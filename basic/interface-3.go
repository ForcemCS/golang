package main

import "fmt"

type Writer interface {
	Write(content string) int
}

type Response struct {
	Body string
}

func (w *Response) Write(content string) int {

	w.Body = content

	return len(content)
}

func myFprint(w Writer, content string) int {

	n := w.Write(content)

	return n
}

func handler(w Writer) {

	myFprint(w, "Hello World")

}

func main() {

	response := Response{
		Body: "",
	}

	handler(&response)

	fmt.Println(response.Body)
}
