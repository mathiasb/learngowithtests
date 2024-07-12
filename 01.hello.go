package main

import "fmt"

func Hello(s string) string {
	if s != "" {
		return "Hello, " + s + "!"
	}
	return "Hello, World!"
}

func main() {
	fmt.Println(Hello(""))
}
