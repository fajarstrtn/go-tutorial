// Comments are non-executable text used to explain code logic-
// or temporarily disable specific sections of a program.
// The Go compiler completely ignores these annotations during execution.
//
// Go supports two primary styles of comments:
// 1. Single-Line Comments: These start with two forward slashes (//).
// 2. Multi-Line Comments : These start with /* and end with */. They can span across multiple lines.
package comment

import "fmt"

// This is a single-line comment.
func PutSingleLineComment() {
	fmt.Println("Hello World") // Output: Hello World
}

/*
 * This is a multi-line comment. */
func PutMultiLineComment() {
	fmt.Println("Hello World") // Output: Hello World
}
