package identifier

import "fmt"

// Identifiers are the user-defined names of the program components.
// In Go, an identifier can be a variable name, function name,
// constant, statement label, package name, or type.
//
// A variable can have a short name (like x and y)
// or a more descriptive name (e.g., age, price, carName).
//
// There are nine identifiers available in the code:
// 1. identifier      : Name of a package
// 2. Identify		  : Name of a function
// 3. name			  : Name of a variable
// 4. _nickName		  : Name of a variable
// 5. Name			  : Name of a variable
// 6. nickN4me		  : Name of a variable
// 7. name2			  : Name of a variable
// 8. full_name       : Name of a variable
// 9. ExportedVariable: Name of an exported variable
//
// There are certain valid rules for defining a valid Go identifier.
// These rules should be followed, otherwise, we will get a compile-time error.
// The name of the identifier:
// 1. Must begin with a letter or an underscore(_)
// 2. May contain the letters 'a-z' or 'A-Z' or digits 0-9 as well as the character '_'
// 3. Should not start with a digit
// 4. Is case-sensitive
// 5. Doesn't allow keyword to become an identifier name
// 6. has no limit on the length of the name of the identifier,
// but it is advisable to use an optimum length of 4–15 letters only
// 7. Cannot contain spaces
//
// Variable name with more than one word can be difficult to read.
// There are several techniques you can use to make them more readable:
// 1. Camel Case  : Each word, except the first, starts with a capital letter (e.g., fullName, graduatedSince)
// 2. Pascal Case : Each word starts with a capital letter (e.g., FullName, GraduatedSince)
// 3. Snake Case  : Each word is separated by an underscore character, but it's uncommon (e.g., full_name, graduated_since)
//
// Camel Case is the standard convention in Go.
// The first letter of the identifier is lowercase, and the first letter of subsequent-
// words is uppercase (e.g., userID, parseRequest)
//
// for most identifiers, including private (unexported) variables, functions, and methods.
func Identify() {
	var name string = "Emerson Santiago"
	var _nickName string = "Emerson"
	var Name string = "Lillie Moon"
	var nickN4me string = "Moon"
	var name2 string = "Davis"
	var full_name string = "Pearl Davis"

	fmt.Println(name)                 // Output: Emerson Santiago
	fmt.Println(_nickName)            // Output: Emerson
	fmt.Println(Name)                 // Output: Lillie Moon
	fmt.Println(nickN4me)             // Output: Moon
	fmt.Println(name2)                // Output: Davis
	fmt.Println(full_name)            // Output: Pearl Davis
	fmt.Println(ExportedVariable)     // Output: Exported Variable
	fmt.Println(packageScopeVariable) // Output: Package-Scope Variable
	fmt.Println(carName)              // Output: Toyota Corolla
	fmt.Println(carType)              // Output: Sedan
	fmt.Println(cc)                   // Output: 1600
}
