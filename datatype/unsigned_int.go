package datatype

import (
	"fmt"
	"math"
	"unsafe"
)

func GetUnsignedIntegers() {
	// An unsigned integer (or uint) is an integer type that cannot hold negative numbers.
	// It only stores zero and positive values.
	//
	// Go categorizes types into basic, aggregate (arrays, structs),
	// reference (pointers, slices, maps, channels), and interface types.
	//
	// Numeric Types (Integers):
	// | Type				| Size			| Range											|
	// | uint8 (alias byte)	| 8 bits		| 0 to 255										|
	// | uint16				| 16 bits		| 0 to 65535									|
	// | uint32				| 32 bits		| 0 to 4,294,967,295 							|
	// | uint64				| 64 bits		| 0 to 18,446,744,073,709,551,615				|
	// | uint				| Platform dep	| Usually 32 or 64 bits. Use this by default	|
	var a uint8 = 8
	fmt.Printf("%d\n", a)                // Output: 8
	fmt.Printf("%T\n", a)                // Output: uint8
	fmt.Printf("%v\n", unsafe.Sizeof(a)) // Output: 1

	var b uint16 = 16_000
	fmt.Printf("%d\n", b)                // Output: 16000
	fmt.Printf("%T\n", b)                // Output: uint16
	fmt.Printf("%v\n", unsafe.Sizeof(b)) // Output: 2

	var c uint32 = 32_000_000
	fmt.Printf("%d\n", c)                // Output: 32000000
	fmt.Printf("%T\n", c)                // Output: uint32
	fmt.Printf("%v\n", unsafe.Sizeof(c)) // Output: 4

	var d uint64 = 64_000_000_000
	fmt.Printf("%d\n", d)                // Output: 64000000000
	fmt.Printf("%T\n", d)                // Output: uint64
	fmt.Printf("%v\n", unsafe.Sizeof(d)) // Output: 8

	var e uint8 = 210
	var f uint16 = uint16(e)

	fmt.Printf("%d\n", e)                // Output: 210
	fmt.Printf("%T\n", e)                // Output: uint8
	fmt.Printf("%v\n", unsafe.Sizeof(e)) // Output: 1

	fmt.Printf("%d\n", f)                // Output: 210
	fmt.Printf("%T\n", f)                // Output: uint16
	fmt.Printf("%v\n", unsafe.Sizeof(f)) // Output: 2

	// Because it doesn't need to waste a bit storing-
	// the positive or negative sign (the Sign Bit),
	// it can store positive numbers twice-
	// as large as its signed counterpart (int).
	var g, h uint = 11_200_000, 81_825_000

	fmt.Printf("%d\n", g)                // Output: 11200000
	fmt.Printf("%T\n", g)                // Output: uint
	fmt.Printf("%v\n", unsafe.Sizeof(g)) // Output: 8

	fmt.Printf("%d\n", h)                // Output: 81825000
	fmt.Printf("%T\n", h)                // Output: uint
	fmt.Printf("%v\n", unsafe.Sizeof(h)) // Output: 8

	fmt.Printf("g + h = %v\n", (g + h))  // Output: g + h = 93025000
	fmt.Printf("g - h = %v\n", (g - h))  // Output: g - h = 18446744073638926616
	fmt.Printf("g * h = %v\n", (g * h))  // Output: g * h = 916440000000000
	fmt.Printf("g / h = %v\n", (g / h))  // Output: g / h = 0
	fmt.Printf("g %% h = %v\n", (g % h)) // Output: g % h = 11200000

	// If a uint hits 0 and you subtract 1 from it, it doesn't become -1.
	// It quietly wraps around to its absolute maximum value.
	// This is one of the most common sources of bugs in Go programs.
	var i uint8 = math.MaxUint8
	var j uint8 = i + 1

	fmt.Printf("%d\n", j) // Output: 0
}
