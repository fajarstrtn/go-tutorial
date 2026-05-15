package format

import "fmt"

func PrintSomething() {
	message1, message2 := "Hello World", "John Doe"

	// The Print() function prints its arguments with their default format.
	// It does not automatically add a newline (\n) at the end.
	// It only adds spaces between arguments if neither argument is a string.
	//
	// This function is used when you just want to dump data-
	// to the console without complex formatting rules.
	//
	// When to use Print() function:
	// 1. You want full control over spacing and line breaks
	// 2. Low-level or very simple output
	fmt.Print(message1)
	fmt.Print(message2)

	// You must manage spaces and newlines yourself.
	fmt.Print(message1 + " ")
	fmt.Print(message2 + "\n")

	// If you want to print the arguments in new lines,
	// you need to use \n (it creates new lines).
	fmt.Print(message1, "\n")
	fmt.Print(message2, "\n")

	// It's also possible to only use one Print() function for printing multiple variables.
	fmt.Print(message1, "\n", message2)

	// If you want to add a space between string arguments, use space.
	fmt.Print(message1, " ", message2, "\n")

	// Print() function inserts a space between the arguments if neither are strings.
	fmt.Print(10, 20, "\n")     // Output: 10 20
	fmt.Print(10.9, true, "\n") // Output: 10.9 true
}
