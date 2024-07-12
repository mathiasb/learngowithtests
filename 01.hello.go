package main

import "fmt"

const englishHelloPrefix = "Hello"

func Hello(s string) string {
	if s == "" {
		s = "World"
	}
	return englishHelloPrefix + ", " + s + "!"
}

func main() {
	fmt.Println(Hello(""))
}
