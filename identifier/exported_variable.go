package identifier

// Identifier which is allowed to access it-
// from another package is known as the exported identifier.
// An identifier (a variable, function, type, constant, or struct field)-
// is exported when it is visible and accessible to packages-
// other than the one in which it is defined.
//
// Exported Variable  : Starts with an uppercase letter (e.g., Config, Calculate)
// Unexported Variable: Starts with a lowercase letter (e.g., secretKey, helperFunc). These are "package-private."
//
// The Go compiler enforces visibility at the package level.
// 1. Scope            : If an identifier is unexported, it can only be seen-
// by other files within the same package.
// 2. Naming Convention: This approach eliminates "keyword noise" in the code.
// You can tell at a glance whether a function is part of a package's public API-
// or an internal helper just by looking at its name.
// 3. Struct Fields    : This rule also applies to fields within a struct.
// If a struct is exported but its fields are lowercase,
// those fields cannot be accessed or modified from outside the package.
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
//
// Expert Go developers often use unexported fields to create Opaque Structs.
// This ensures that the internal state of a component cannot be "clobbered"-
// by an external user, maintaining a strict contract through exported methods.
//
// To use an exported identifier, you must:
// 1. Define it in a package with an uppercase starting letter
// 2. Import that package into your current file
// 3. Access it using the package name as a prefix (e.g., pkgName.Identifier)
var (
	ExportedVariable     string = "Exported Variable"
	packageScopeVariable string = "Package-Scope Variable"
)
