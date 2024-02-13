package main

// parentheses mean it's a factored import statement, better style than individual imports
import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand" // random library
)

// name of parameter then type, return type after these
// type of first can be omitted if they're of the same type
func add(x, y int) int {
	return x + y
}

// functions can return multiple types, listed afterwards in parentheses
func swap(x, y string) (string, string) {
	return y, x
}

// return values can be named and treated as vars
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // "naked" return; just returns the declared return vars. only good for short funcs
}

// var used to declare a list of variables, with the type last
// bools automatically declared as false
// var can also give them all values, as seen below
var c, java, python bool
var one, two, three = 1, 2, 3

// shows the types of vars in go
func types() {
	// int int8 int16 int32 int64, along with unsigned versions called uint and uintptr
	// byte is alias for uint8, rune alias for int32
	// float32 float64
	// complex64 complex128
	var (
		maxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Println(maxInt, "\n", z)
}

// undeclared vars given their zero value
// strings -> ""
// ints    -> 0
// bools   -> false

// var typing is converted by equation T(v), where T is new type v is old
func main() {
	fmt.Println("My favorite number is:", rand.Intn(10))

	// format string printing similar to C
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	fmt.Println(math.Pi) // has to be capital since imported

	// := declares variable type as well. can't be used outside a function
	// and can't be used twice on the same variable
	// can be used to define multiple at one
	// infers var typing based on what it's set to
	x, y, z := "hello", "there", "world"
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(x, y, z)

	fmt.Println(split(17)) // returns 7 and 10

	fmt.Println(c, java, python)
	fmt.Println(one, two, three)

	types()

	// constants given the const keyword
	const Constant = "I am constant."

	fmt.Println("Our constant is", Constant)

}
