package main

import "fmt"

func main() {
	a, b := "cat", "dog"
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
}