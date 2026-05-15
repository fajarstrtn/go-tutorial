package format

import "log"

func PrintSomethingWithLog() {
	// For real apps:
	// 1. Adds timestamp
	// 2. Writes to stderr
	// 3. Better for production
	logMessage := "User has been created"
	log.Println(logMessage) // Output: 2026/05/16 01:29:12 User has been created
}
