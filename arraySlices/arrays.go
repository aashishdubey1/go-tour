package arrayslices

import (
	"fmt"
	"strings"
)

func Run() {

	fmt.Println("------Arrays------")

	// The type [n]T is an array of n values of type T.
	// An array's length is part of its type, so arrays cannot be resized.

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var b [10]int
	fmt.Println(b)

	fmt.Println("------Slices------")
	// An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.
	// The type []T is a slice with elements of type T.
	var s []int = primes[1:4]
	fmt.Println(s)
	// A slice is formed by specifying two indices, a low and high bound, separated by a colon: excludes the last one.

	// Slices are like references to arrays
	// so slice basically consist of parts. 1.Length 2.Capacity 3.Pointer
	//-> length = length of the slice that we get from startIndex and endIndex
	//-> capacity = length of arr from startIndex of slice and end of arr
	//-> pointer = pointer to startIndex of slice in array

	// A slice does not store any data, it just describes a section of an underlying array.
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	// Other slices that share the same underlying array will see those changes.

	fmt.Println("------Slice literals------")
	// A slice literal is like an array literal without the length.
	// This is an array literal:
	// [3]bool{true, true, false}
	// And this creates the same array as above, then builds a slice that references it:
	// []bool{true, true, false}

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	p := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(p)

	q = q[1:4]
	fmt.Println(q)

	q = q[:2]
	fmt.Println(q)

	q = q[1:]
	fmt.Println(q)

	printSlice(s)

	fmt.Println("------Nill Slice------")
	// The zero value of a slice is nil.
	// A nil slice has a length and capacity of 0 and has no underlying array.
	var k []int
	fmt.Println(k, len(k), cap(k))
	if k == nil {
		fmt.Println("nil!")
	}

	fmt.Println("------Slice with make------")
	// make is a built-in function used to initialize and allocate memory for reference-like types.
	// Specifically make is used for :- Slices, maps, channels
	// slices need underlying memory, and make allocates + initializes that memory.

	// WHY WE USE MAKE TO CREATE SLICES ?
	// We use make when we want a slice with a predefined length
	// so that we can safely access elements by index.
	//
	// A slice literal like []int{} creates an empty slice
	// with length 0, so accessing x[0] causes "index out of range" panic.
	//
	// make allocates underlying memory and sets the slice length,
	// so indexing is safe up to len-1.

	// if we try to do this , result will be panic
	// x := []int{}
	// fmt.Println(x)
	// x[0] = 10
	// fmt.Println(x)

	// can do with make
	y := make([]int, 5, 10) // it inatialize array with int default values = 0
	fmt.Println(y)
	y[0] = 20
	fmt.Println(y)
	// WHY APPEND WORKS THEN ?
	// because earlier we are inserting a element at index that is not present in that array
	// what append does is check for cap if it's full allocate a new array copy existing el , add new one and return new slice
	// if len(s) + needed <= cap(s):
	//      place element(s) in existing array
	//      return updated slice header
	//  else:
	//      allocate new array with larger capacity
	//      copy old elements
	//      add new element(s)
	//      return new slice
	// append allocates a new array when capacity is exceeded; otherwise it reuses the existing array and just grows length.

	fmt.Println("------Slices of slices------")
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	fmt.Println("------Appending to a Slice------")
	// func append(s []T, vs ...T) []T
	// The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.

	var g []int
	printSlice(s)

	// append works on nil slices.
	g = append(g, 0)
	printSlice(g)

	// The slice grows as needed.
	g = append(g, 1)
	printSlice(g)

	// We can add more than one element at a time.
	g = append(g, 2, 3, 4)
	printSlice(g)

	l := make([]int, 3, 3)

	l = append(l, 5, 5, 5)
	printSlice(l)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
