// In Go, every file must belong to a package.
// Think of a package as a folder of related code.
//
// The main package tells the Go compiler that this file is intended-
// to be an executable program rather than a library (a collection of tools for other programs).
// Also, this is special because this program is meant to be run, not imported.
// The main package means "this is the starting point of my app".
//
// If you want to build an executable (e.g., go run, go build), the package must be main.
// If you change this to package helper,
// you wouldn't be able to run the file directly as a program,
// it would just be a package of code for others to use.
//
// The below explanation is a full execution flow of Go:
// 1. Go starts the program
// 2. Finds package main
// 3. Looks for func main()
// 4. Runs code inside main()
// 5. Prints "Hello World"
// 6. Program ends
//
// There are concerns you need to know:
// 1. No semicolons needed (Go inserts them automatically)
// 2. Imports must be used; unused imports cause errors
// 3. File must be named something.go
package main

// The import keyword is used to bring in code from other packages.
import (
	"fmt"

	"github.com/fajarstrtn/go-tutorial/comment"
	"github.com/fajarstrtn/go-tutorial/format"
	"github.com/fajarstrtn/go-tutorial/identifier"
	"github.com/fajarstrtn/go-tutorial/introduction"
)

// When you run a program, Go automatically starts executing main function.
// No main function means nothing runs.
//
// The func keyword is how you declare a function.
// The main() function is a special function name.
// In an executable program, the main function is the entry point.
// When you run your program, the computer starts executing code-
// inside the curly braces {} of this function.
//
// In Go, the opening curly brace { must be on the same line as the function declaration.
// Moving it to the next line will cause a syntax error.
func main() {
	introduction.Greet()
	comment.PutSingleLineComment()
	comment.PutMultiLineComment()
	format.PrintSomething()
	format.PrintSomethingWithNewLine()
	format.PrintSomethingWithLog()
	format.PrintSomethingWithSprintf()
	format.PrintSomethingWithFormattingVerbs()

	exportVariable()
}

func exportVariable() {
	fmt.Printf("%s is being called\n", identifier.ExportedVariable) // Output: Exported Variable is being called
}
