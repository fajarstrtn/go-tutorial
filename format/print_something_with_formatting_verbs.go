package format

import "fmt"

func PrintSomethingWithFormattingVerbs() {
	x, y, z, b := 10, 3.141, "Hello World", false

	// The Printf() function first formats its argument based on-
	// the given formatting verb and then prints them.
	// It gives you precise control over how data looks-
	// and does not add a newline automatically.
	//
	// Go offers several formatting verbs that can be used with the Printf() function.
	// Format verbs start with %a and it prints using a format string.
	//
	// The following verbs can be used with all data types:
	// 1. %v : Prints the value in the default format
	// 2. %+v: Prints the value in its default format with the field names-
	// included when dealing specifically with structs
	// 3. %#v: Prints the value in Go-syntax format
	// 4. %T : Prints the type of the value
	// 5. %% : Prints the % sign
	//
	// The + or # symbol is a flag that modifies the behavior of the general %v verb.
	fmt.Printf("%v\n", x)   // Output: 10
	fmt.Printf("%#v\n", x)  // Output: 10
	fmt.Printf("%T\n", x)   // Output: int
	fmt.Printf("%v%%\n", x) // Output: 10%

	fmt.Printf("%v\n", y)   // Output: 3.141
	fmt.Printf("%#v\n", y)  // Output: 3.141
	fmt.Printf("%T\n", y)   // Output: float64
	fmt.Printf("%v%%\n", y) // Output: 3.141%

	fmt.Printf("%v\n", z)   // Output: Hello World
	fmt.Printf("%#v\n", z)  // Output: "Hello World"
	fmt.Printf("%T\n", z)   // Output: string
	fmt.Printf("%v%%\n", z) // Output: Hello World%

	fmt.Printf("%v\n", b)   // Output: false
	fmt.Printf("%#v\n", b)  // Output: false
	fmt.Printf("%T\n", b)   // Output: bool
	fmt.Printf("%v%%\n", b) // Output: false%

	type User struct {
		Name string
		Age  int
	}

	user := User{"John Doe", 20}

	// Using %v is the same as "use Go's default rule."
	fmt.Printf("%v\n", user) // Output: {John Doe 20}

	// It's Go-syntax representation (great for debugging).
	fmt.Printf("%#v\n", user) // Output: format.User{Name:"John Doe", Age:20}

	// The primary use of %+v is to display the field names-
	// along with their corresponding values,
	// which is extremely useful for debugging.
	//
	// For basic types (e.g., integers, strings, etc.),
	// the + flag generally has no effect,
	// and the output is the same as %v.
	//
	// Use %+v while learning structs.
	fmt.Printf("%+v\n", user) // Output: {Name:John Doe Age:20}

	// It works with any type and makes %v perfect for debugging.
	// It works automatically.
	// No loops needed.
	fmt.Printf("%v\n", []int{1, 2, 3})         // Output: [1 2 3]
	fmt.Printf("%v\n", map[string]int{"a": 1}) // Output: map[a:1]

	// The following verbs can be used with the integer data type:
	// 1. %b   : Base 2
	// 2. %d   : base 10
	// 3. %+d  : Base 10 and always show sign
	// 4. %o   : Base 8
	// 5. %O   : Base 8 with leading 0o
	// 6. %x   : Base 16 lowercase
	// 7. %X   : Base 16 uppercase
	// 8. %#x  : Base 16 with leading 0x
	// 9. %4d  : Pad with spaces (width 4, right justified)
	// 10. %-4d: Pad with spaces (width 4, left justified)
	// 11. %04d: Pad with zeroes (width 4)
	fmt.Printf("%b\n", x)   // Output: 1010
	fmt.Printf("%d\n", x)   // Output: 10
	fmt.Printf("%+d\n", x)  // Output: +10
	fmt.Printf("%o\n", x)   // Output: 12
	fmt.Printf("%O\n", x)   // Output: 0o12
	fmt.Printf("%x\n", x)   // Output: a
	fmt.Printf("%X\n", x)   // Output: A
	fmt.Printf("%#X\n", x)  // Output: 0XA
	fmt.Printf("%5d\n", x)  // Output:    10
	fmt.Printf("%-5d\n", x) // Output: 10
	fmt.Printf("%05d\n", x) // Output: 00010

	// The following verbs can be used with the float data type:
	// 1. %e   : Scientific notation with 'e' as exponent
	// 2. %f   : Decimal point, no exponent
	// 3. %.2f : Default width, precision 2
	// 4. %6.2f: Width 6, precision 2
	// 5. %g   : Exponent as needed, only necessary digits
	fmt.Printf("%e\n", y)     // Output: 3.141000e+00
	fmt.Printf("%f\n", y)     // Output: 3.141000
	fmt.Printf("%.1f\n", y)   // Output: 3.1
	fmt.Printf("%10.2f\n", y) // Output:       3.14
	fmt.Printf("%g\n", y)     // Output: 3.141

	// The following verbs can be used with the string data type:
	// 1. %s  : Prints the value as plain string
	// 2. %q  : Prints the value as a double-quoted string
	// 3. %8s : Prints the value as plain string (width 8, right justified)
	// 4. %-8s: Prints the value as plain string (width 8, left justified)
	// 5. %x  : Prints the value as hex dump of byte values with lowercase
	// 6. %X  : Prints the value as hex dump of byte values with uppercase
	// 7. % x : Prints the value as hex dump with spaces with lowercase
	// 8. % X : Prints the value as hex dump with spaces with uppercase
	fmt.Printf("%s\n", z)    // Output: Hello World
	fmt.Printf("%q\n", z)    // Output: "Hello World"
	fmt.Printf("%15s\n", z)  // Output:     Hello World
	fmt.Printf("%-15s\n", z) // Output: Hello World
	fmt.Printf("%x\n", z)    // Output: 48656c6c6f20576f726c64
	fmt.Printf("%X\n", z)    // Output: 48656C6C6F20576F726C64
	fmt.Printf("% x\n", z)   // Output: 48 65 6c 6c 6f 20 57 6f 72 6c 64
	fmt.Printf("% X\n", z)   // Output: 48 65 6C 6C 6F 20 57 6F 72 6C 64

	// The following verb can be used with the boolean data type:
	// 1. %t: Value of the boolean operator in true or false format (same as using %v)
	fmt.Printf("%t\n", b) // Output: false

	// Think of fmt as Go's text formatting engine.
	// Main printing functions you'll use:
	// 1. fmt.Print* : Prints to stdout (terminal), Returns bytes written
	// 2. fmt.Fprint*: Prints any io.Writer, Returns bytes written
	// 3. fmt.Sprint*: Prints to string, Returns string
	//
	// Rule of thumb:
	// 1. Print  : rare, manual control
	// 2. Println: learning & debugging
	// 3. Printf : structured output
	//
	// Golden rules to memorize:
	// 1. Println: Prints daily debugging
	// 2. Printf : Prints formatted output
	// 3. Sprintf: Prints formatted string
	// 4. %v     : Prints value
	// 5. %T     : Prints type
	// 6. %+v    : Prints struct fields
	// 7. %#v    : Prints Go-syntax
	fmt.Print("XYZ\n")              // Output: XYZ
	fmt.Println("X", "Y")           // Output: X Y
	fmt.Printf("%s %d\n", "AB", 10) // Output: AB 10
}
