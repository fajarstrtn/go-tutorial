package datatype

import (
	"fmt"
	"unsafe"
)

func GetSignedIntegers() {
	// Go categorizes types into basic, aggregate (arrays, structs),
	// reference (pointers, slices, maps, channels), and interface types.
	//
	// Numeric Types (Integers):
	// | Type				| Size			| Range														|
	// | int8				| 8 bits		| -128 to 127												|
	// | int16				| 16 bits		| -32768 to 32767											|
	// | int32 (alias rune) | 32 bits		| -2,147,483,648 to 2,147,483,647 							|
	// | int64				| 64 bits		| -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807	|
	// | int				| Platform dep	| Usually 32 or 64 bits. Use this by default				|
	var a int8 = 8
	fmt.Printf("%d\n", a)                // Output: 8
	fmt.Printf("%T\n", a)                // Output: int8
	fmt.Printf("%v\n", unsafe.Sizeof(a)) // Output: 1

	var b int16 = 16_000
	fmt.Printf("%d\n", b)                // Output: 16000
	fmt.Printf("%T\n", b)                // Output: int16
	fmt.Printf("%v\n", unsafe.Sizeof(b)) // Output: 2

	var c int32 = 32_000_000
	fmt.Printf("%d\n", c)                // Output: 32000000
	fmt.Printf("%T\n", c)                // Output: int32
	fmt.Printf("%v\n", unsafe.Sizeof(c)) // Output: 4

	var d int64 = 64_000_000_000
	fmt.Printf("%d\n", d)                // Output: 64000000000
	fmt.Printf("%T\n", d)                // Output: int64
	fmt.Printf("%v\n", unsafe.Sizeof(d)) // Output: 8

	// An arithmetic operation between an int8 and an int16
	// (or any different data types including int and int32) variable
	// causes a compilation error because Go is a statically
	// and strictly typed language that does not support implicit type conversion.
	//
	// Go requires the operands of an operation-
	// (e.g., addition, comparison, or assignment)
	// to have the exact same type.
	//
	// This design choice is intentional to avoid potential errors
	// and unexpected behavior that can arise-
	// from automatic (implicit) type conversions.
	//
	// int8 and int16 are fundamentally different types-
	// with different memory sizes and ranges of values.
	//
	// Go compiler does not automatically "widen" the int8-
	// to an int16 before performing the calculation,
	// even though it would be a "safe" operation
	// in terms of not losing data.
	//
	// To perform addition, you must explicitly-
	// convert one of the variables to match the type of the other.
	// The result of the operation will then have the new, unified type.
	var e int8 = 120
	var f int16 = 16_000
	var g int16 = int16(e)

	fmt.Printf("%d\n", e) // Output: 120
	fmt.Printf("%d\n", f) // Output: 16000
	fmt.Printf("%d\n", g) // Output: 120

	// You can use %v to print the default format.
	fmt.Printf("f + g = %v\n", (f + g))  // Output: f + g = 16120
	fmt.Printf("f - g = %v\n", (f - g))  // Output: f - g = 15880
	fmt.Printf("f * g = %v\n", (f * g))  // Output: f * g = 19456
	fmt.Printf("f / g = %v\n", (f / g))  // Output: f / g = 133
	fmt.Printf("f %% g = %v\n", (f % g)) // Output: f % g = 40

	// If you do not specify a type, the default type will be int.
	// int is the default numeric types used by the compiler-
	// when it infers types from untyped constants.
	//
	// When you declare a variable using short-variable declaration (:=)
	// and provide a numeric literal without a type,
	// Go assigns a default type based on the format of the literal,
	// where integer literals set to default to int.
	//
	// The size of an int is platform-dependent.
	// It's usually 32 bits on 32-bit systems and 64 bits on 64-bit systems.
	h, i := 450_000, 750_000

	fmt.Printf("%d\n", h)                // Output: 450000
	fmt.Printf("%T\n", h)                // Output: int
	fmt.Printf("%v\n", unsafe.Sizeof(h)) // Output: 8

	fmt.Printf("%d\n", i)                // Output: 750000
	fmt.Printf("%T\n", i)                // Output: int
	fmt.Printf("%v\n", unsafe.Sizeof(i)) // Output: 8

	fmt.Printf("h + i = %v\n", (h + i))  // Output: h + i = 1200000
	fmt.Printf("h - i = %v\n", (h - i))  // Output: h - i = -300000
	fmt.Printf("h * i = %v\n", (h * i))  // Output: h * i = 337500000000
	fmt.Printf("h / i = %v\n", (h / i))  // Output: h / i = 0
	fmt.Printf("h %% i = %v\n", (h % i)) // Output: h % i = 450000

	// When you declare an int64 variable-
	// in a Go program compiled for a 32-bit system,
	// the variable will still be a 64-bit integer,
	//  but operations on it will be less performant than on a 64-bit system.
	//
	// int64 is an explicitly sized type in Go, meaning it is always 64 bits wide-
	// regardless of the underlying architecture (unlike the generic int type,
	// which is 32 bits on a 32-bit system).
	// Go specification defines all numeric types-
	// with fixed sizes, ensuring portability of size.
	//
	// Since a 32-bit processor's Arithmetic Logic Unit (ALU)
	// can only handle 32-bit operations natively,
	// Go compiler generates extra instructions to perform int64 operations
	// (like addition, multiplication, and division) using-
	// multiple 32-bit operations.
	// This process is known as software emulation.
	//
	// The primary consequence is a performance penalty.
	// Operations on int64 variables will be slower compared to-
	// operations on the native int32 type on that same 32-bit system,
	// and significantly slower than on a 64-bit system-
	// where int64 operations are single instructions.
	//
	// An int64 variable will occupy 8 bytes of memory-
	// even on a 32-bit system, potentially increasing memory usage-
	// compared to using int32 if smaller values suffice.
	//
	// The code will work correctly across architectures,
	// but for performance-critical sections on 32-bit systems,
	// it's often recommended to use int or int32-
	// if the values will fit within a 32-bit range.
	//
	// Declaring an int64 on a 32-bit system works just fine,
	// as the Go compiler handles the underlying logic,
	// but it results in a performance trade-off due to-
	// the need for 64-bit arithmetic to be emulated in software.
	var j int32 = 7_000_000
	var k int64 = 1_700_000_000

	fmt.Printf("%d\n", j)                // Output: 7000000
	fmt.Printf("%T\n", j)                // Output: int32
	fmt.Printf("%v\n", unsafe.Sizeof(j)) // Output: 4

	fmt.Printf("%d\n", k)                // Output: 1700000000
	fmt.Printf("%T\n", k)                // Output: int64
	fmt.Printf("%v\n", unsafe.Sizeof(k)) // Output: 8
}
