package identifier

// Identifier which is allowed to access it-
// from another package is known as the exported identifier.
// An identifier (a variable, function, type, constant, or struct field)-
// is exported when it is visible and accessible to packages-
// other than the one in which it is defined.
//
// Exported Variable  : Starts with an uppercase letter (e.g., Config, Calculate)
// Unexported Variable: Starts with a lowercase letter (e.g., secretKey, helperFunc)
// These are "package-private."
//
// In Go, the concept of an Exported Identifier-
// is the language's built-in mechanism for access control.
// Unlike languages that use keywords like public or private,
// Go uses a simple, elegant rule based on the casing of the first letter.
//
// The exported identifiers are those identifiers-
// which obey the following condition:
// 1. The first character of the exported identifier's name-
// should be in the Unicode upper case letter.
// 2. The identifier should be declared in the package block-
// or be a variable, function, type, or method name within that package.
//
// The uniqueness of the identifiers means the identifier is unique-
// from the other set of the identifiers available in your program,
// or in the package and they are not exported.
//
// The ExportedVariable is recognized and accessible in its package.
//
// Go uses Pascal Case (also known as Upper Camel Case) for identifiers-
// that need to be visible or accessible from other packages (exported).
// If an identifier starts with an uppercase letter, it is exported.
// It is used for public variables, functions, methods,
// struct names, and interface names.
//
// Snake Case (using underscores) is generally discouraged in Go.
// The only exception is for file names, where snake_case.go is acceptable,
// although camelCase.go is also common.
var ExportedVariable string = "Exported Variable"
