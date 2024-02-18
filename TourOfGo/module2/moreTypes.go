package main

import (
	"fmt"
	"math"
	"strings"
)

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

	// structs can also be defined to a variable, first class object
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{5, false},
	}
	fmt.Println(s)

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

	// taking slices from arrays as well
	s := primes[1:4] // EXCLUSIVE NOT INCLUSIVE; so this selects elements 1-3
	fmt.Println(s)
	// slices don't store any data, but describes the sections of an underlying array
	// so if we change an element in s, it'll be reflected in primes
	s[2] = 17
	fmt.Println(primes)

	// slice literals is an array literal w/o the length
	myList := []bool{true, false, true} // makes the same array, and the slice just references it
	fmt.Println(myList)

	// slices have defaults
	//var a [10]int = a[:10], a[0:10], a[:], a[0:] all the same

	// slice definitions:
	// length: number of elements it contains; obtained with len(s)
	// capacity: number of elements in the underlying array COUNTING FROM FIRST ELEMNT
	//   	     IN THE SLICE; obtained with cap(s)
	fmt.Println("Length", len(s), "\nCapacity", cap(s))

	// zero value of a slice is 'nil'; length and capacity of 0
	var empty []int
	fmt.Println("Empty slice:", empty, len(empty), cap(empty))

	// using make will create dynamically-sized arrays
	dynamic := make([]int, 4, 5) // type, length, capacity
	fmt.Println(dynamic)         // fills with 0's for zero value

	// 2d slices! slices can have any value including slices
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	for i := 0; i < len(board); i++ {
		// print the board, joining each underlying string together with the join method
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// can use append on slices
	// append(s []T, vs ...T) []T
	// first parameter is s slice of type T, with the rest being values of type T to append
	// results in another slice with original slice + appended
	dynamic = append(dynamic, 8, 1)
	fmt.Println(dynamic, "\nNew capacity:", cap(dynamic))
	// increases the capacity if new array exceeds bounds
}

func ranges() {
	fmt.Println("\nBEGIN RANGES -----------")
	defer fmt.Println("END RANGES -----------")

	// range can iterate a slice or map
	pow := []int{1, 2, 4, 8, 16, 32, 64}
	for i, v := range pow {
		// i = index, v = element at index
		fmt.Println("Index:", i, " Element:", v)
	}

	// index or value can be omitted from the range
	for _, v := range pow {
		fmt.Println("Value:", v)
	}
}

func maps() {
	fmt.Println("\nBEGIN MAPS -----------")
	defer fmt.Println("\nEND MAPS -----------")
	// maps key-value pairs
	// zero value is nil, and a nil map has no keys and keys can't be added
	m := make(map[string]string)
	m["Hello"] = "world!"
	fmt.Println(m)
	fmt.Println(m["Hello"])

	// like struct literals but just require keys
	// if top level types is just a type name, can omit it from elements
	type vertex struct {
		x, y int
	}
	var exam = map[string]vertex{
		"example":     {123, 345},
		"of omitting": {456, 567},
	}
	fmt.Println(exam)

	// deleting an elem from a map
	delete(exam, "example")
	fmt.Println(exam)

	// test to see if key is present in value
	test, yes := m["Hello"]
	fmt.Println("test:", test, "\nis it in map?", yes)
	// test will be the element there or zero value, yes is t/f of whether present
}

// functions cannot be declared in one another, so we explain functions in global space
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// functions may be closers, meaning it's a function value that references variables from
// outside of it's body. can access and assign referenced variables

// calls all the above functions to display how different types and functions work
// in golang
func main() {
	pointers() // see how pointers work
	structs()  // see how structs work
	arrays()   // see how arrays work
	ranges()   // see how ranges work
	maps()     // see how maps work

	fmt.Println("\nBEGIN FUNCTIONS -----------")
	defer fmt.Println("END FUNCTIONS	 -----------")
	calc := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println("computing:", compute(calc))

}
