package introduction

// The fmt (short for format) is a standard library package.
// It contains functions for formatting text, including printing to the console.
// This imports the standard formatting package of Go.
//
// In Go, you cannot import a package and then not use it.
// If you import fmt but don't use it in your code,
// the program will throw an error and refuse to compile.
// This keeps the code clean and the binaries small.
import "fmt"

func Greet() {
	// The fmt package is the package you imported.
	// This tells Go to look inside the fmt package we imported earlier.
	// The Println() function is the function name.
	// It stands for print line, where it prints the text inside-
	// the quotes and then adds a new line at the end-
	// so the next output starts on a fresh row.
	//
	// Notice the "P" is capitalized.
	// In Go, a function is only exported (accessible from outside its package)-
	// if it starts with a capital letter.
	// Since we are using Println outside of the fmt package, it must be capitalized.
	fmt.Println("Hello World") // Output: Hello World
}
