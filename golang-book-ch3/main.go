package main

import "fmt"

func main() {
	// int and float64
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("1.0 + 1.0 =", 1.0+1.0)

	// strings
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1]) // returns e as int representation
	fmt.Println("Hello " + "World")

	// booleans
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}
