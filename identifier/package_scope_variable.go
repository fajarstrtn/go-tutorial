package identifier

import "fmt"

// In Go, the boundary for visibility is the package, not the file.
//
// When you compile your code, the Go compiler essentially groups-
// all the files belonging to the same package together
// and treats them as one giant file.
// Because of this, any variable declared outside a function-
// (at the package level) is visible to every single file-
// within that same package, even if it starts with a lowercase letter.
//
// These are package-scope variables.
var (
	carName = "Toyota Corolla"
	carType = "Sedan"
	cc      = 1600
)

func GetCarDetail() {
	// This is local variable.
	color := "Midnight Black"

	fmt.Println(carName) // Output: Toyota Corolla
	fmt.Println(carType) // Output: Sedan
	fmt.Println(cc)      // Output: 1600
	fmt.Println(color)   // Output: Midnight Black
}
