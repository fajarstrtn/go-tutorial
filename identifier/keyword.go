package identifier

import "fmt"

// Keywords or Reserved Words are the words in a language-
// that are used for some internal process or represent some predefined actions.
// There are total 25 keywords present in the Go.
//
// These words are therefore not allowed to use as an identifier.
// Doing this will result in a compile-time error.
func GetKeywords() {
	keywords := "break, case, chan, const, continue, default, defer, fallthrough, for, func, go, goto, if, import, interface, map, package, range, return, select, struct, switch, type, var"
	fmt.Println(keywords) // Output: break, case, chan, const, continue, default, defer, fallthrough, for, func, go, goto, if, import, interface, map, package, range, return, select, struct, switch, type, var
}
