package main

import "fmt"

func main() {
	var i interface{} = "Hello"
	fmt.Println(i)

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// panic
	f = i.(float64)
	fmt.Println(f)
}