package main

import "fmt"

// tbis method goes through and explains how pointers work
func pointers() {
	fmt.Println("\nBEGIN POINTERS -----------")
	defer fmt.Println("END POINTERS -----------")

	i, j := 42, 2701 // vars to use for our pointers

	p := &i         // pointer at i
	fmt.Println(*p) // dereference and see value of i through p

	*p = 21
	fmt.Println(i) // change i through the pointer

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

// this method goes through and explains how structs work
func structs() {
	fmt.Println("\nBEGIN STRUCTS -----------")
	defer fmt.Println("END STRUCTS -----------")

	// define a vertex to be a struct
	type Vertex struct {
		x, y int
	}
	fmt.Println(Vertex{1, 2})

	// structs just act as objects
	v := Vertex{2, 3}
	fmt.Println(v.y)

	// can access fields of structs through a pointer
	p := &v   // pointer to vertex
	p.x = 1e9 // Go implicitly dereferences p in order to access x; RAAHHH SCIENTIFIC NOTATION
	fmt.Println(v)

	// values will have implicit/default values
	v1 := Vertex{x: 12}
	fmt.Println(v1) // y automatically set to 0
}

// this method explains how arrays function
func arrays() {
	fmt.Println("\nBEGIN ARRAYS -----------")
	defer fmt.Println("END ARRAYS -----------")
	// [n]T is an array of n values with type T
	var a [10]int // array with 10 entries of type int
	a[1] = 2      // indexes 0-9

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
func main() {
	pointers() // see how pointers work
	structs()  // see how structs work
	arrays()   // see how arrays work

}
