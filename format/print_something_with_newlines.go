package format

import "fmt"

func PrintSomethingWithNewLine() {
	message1, message2 := "Hello World", "John Doe"

	// The Println() function prints with a newline at the end-
	// and automatically adds spaces between arguments.
	// It outputs arguments, spaces them out cleanly, and moves to the next line.
	// It converts values to readable text automatically.
	//
	// This function is also used when you just want to dump data-
	// to the console without complex formatting rules.
	// If you're not sure which one to use, use Println() function.
	//
	// When to use Println() function:
	// 1. Most common choice
	// 2. Quick logging
	// 3. Debugging values
	//
	// When you do Println() function:
	// 1. Detects type
	// 2. Applies default formatting
	// 3. Converts to string
	// 4. Writes to stdout
	//
	// How the Professionals Use Them (Best Practices)
	// 1. Never use fmt.Print/Println for Application Logs:
	// In production, standard output (stdout) is typically reserved for-
	// actual program data or CLI outputs.
	// Logging info, warnings, or errors to stdout via fmt-
	// makes it incredibly hard to separate data from system diagnostics.
	// Pros always use a logger for system events.
	// 2. Leverage fmt.Sprintf for Dynamic String Building:
	// Pros use Sprintf to build URLs, file paths, and custom error messages.
	// Syntax:
	// err := fmt.Errorf("failed to fetch user %d: %w", userID, actualErr)
	// 3. Debugging Structs with %+v:
	// If a pro wants to inspect a complex state or nested struct quickly-
	// during local debugging, they don't print individual fields.
	// They dump the whole thing using the struct-verb combo:
	// Syntax:
	// fmt.Printf("Current state: %+v\n", userProfile)
	// 4. Directing Logs to Files:
	// Pros rarely leave the log package printing to the console in production.
	// They redirect it to a file or an external log aggregator-
	// right at the application's entry point (func main()).
	// Syntax:
	// file, _ := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// log.SetOutput(file)
	fmt.Println(message1, message2) // Output: Hello World John Doe
	fmt.Println("Hello", "World")   // Output: Hello World
	fmt.Println(10, "Hello", true)  // Output: 10 Hello true
}
