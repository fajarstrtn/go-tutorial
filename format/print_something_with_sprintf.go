package format

import "fmt"

func PrintSomethingWithSprintf() {
	const fullName, nickName string = "John Doe", "John"

	// In Go, the fmt package implements formatted I/O with functions.
	// The fmt.Sprintf() function in Go formats according to-
	// a format specifier and returns the resulting string.
	// It returns a string instead of printing.
	//
	// You will use this for:
	// 1. Logs
	// 2. Errors
	// 3. Building strings
	// 4. APIs
	message := fmt.Sprintf("Hello, %s! You can call me %s", fullName, nickName)
	fmt.Println(message) // Output: Hello, John Doe! You can call me John

	error := fmt.Sprintf("%s", "Something went wrong!")
	fmt.Println(error) // Output: Something went wrong!

	// You can also print with width, alignment, and precision.
	fmt.Printf("|%10s|\n", fullName)  // Output: |  John Doe|
	fmt.Printf("|%-10s|\n", fullName) // Output: |John Doe  |
}
