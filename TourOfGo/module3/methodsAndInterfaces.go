package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

/*
 The below is a block about methods (functions with receivers)
*/

// Abs this is a method is a method due to the receiver argument that comes
// between the func declaration and name of the method
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Abs2 can be written as a function as well, and the only difference is the
// receiver that comes before it
func Abs2(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// can only declare a method with a receiver whose type is defined in the same
// package as the method. You cannot declare a method with a receiver whose type
// is defined in another package (which includes the built-in types such as int).

// Scale can give pointers as a receiver type, meaning they can modify whatever
// they point at. pointer receivers are more used than regular as the function
// is used to update the value that it points at
// again, could be rewritten as a function if need be
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/*
methods can take pointers or values in as receivers and still compiled,
while functions must take a pointer.

There are two reasons to use a pointer receiver:
The first is so that the method can modify the value that its receiver points to.
The second is to avoid copying the value on each method call. This can be more efficient
if the receiver is a large struct, for example.
*/

func main() {
	fmt.Println("Hello world!")
}
