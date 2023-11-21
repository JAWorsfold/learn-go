package main

import "fmt"

const y string = "const"

func main() {
	// strings()

	temp()
}

func strings() {
	// var x string = "Hello World" // or
	var x string
	x = "Hello World"
	fmt.Println(x)
	x = "Goodbye"
	fmt.Println(x)
	f()

	// string concat
	const z string = y + y
	fmt.Println(z)
	x += y
	fmt.Println(x)

	// equality
	fmt.Println(x == y)

	// inference
	s := "string"
	fmt.Println(s)
	// not sure how this works with const?

	// multi assignment
	const (
		a = "1"
		b = "b"
		c = "c"
	)
}

func f() {
	// y = "break" // can't be done gives error:
	// cannot assign to y (neither addressable nor a map index expression)
	fmt.Println(y)
}

func temp() {
	fmt.Print("Enter a temperature: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := (input - 32) * 5/9

	fmt.Println(output)
}
