package tutorial

import (
	"fmt"
	"strconv"
)

// Package level variables:
var global_i float64 = 20.5 // package scoped (starts lowercase)
var Global_i float64 = 20.5 // public  scoped (starts uppercase)

// var blocks: just to help organize
var (
	actorName string = "Michael Davis"
	actorAge  int    = 28
	actorType string = "C"
)

// constant block using iota
const (
	_  = iota // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func line_break() {
	fmt.Println()
}

func Tut_Variables() {
	// Three ways to define variables:
	{
		// 1 - use first line if just want to declare
		var i int
		i = 42

		// 2
		var j int = 42

		// 3
		k := 42

		fmt.Println(i, j, k)
	}

	line_break()

	// Printing formatting and getting var type:
	{
		var i int = 42
		fmt.Printf("%v, %T\n", i, i) // %v = value?, %T = type
	}

	line_break()

	// Global tests:
	{
		// shadow variables: using a global name for a local variable
		fmt.Printf("Global I BC: %v, %T\n", global_i, global_i)
		global_i := 25
		fmt.Printf("Global I AD: %v, %T\n", global_i, global_i)
	}

	line_break()

	// Conversion tests:
	{
		var i int = 42
		fmt.Printf("Start val: %v, %T\n", i, i)
		var j float64 = float64(i)
		fmt.Printf("Converted val: %v, %T\n", j, j)

		var k string = string(i) //gets unicode char
		fmt.Printf("Converted str (as unicode): %v, %T\n", k, k)
		k = strconv.Itoa(i) // prints as string
		fmt.Printf("Converted str (normal): %v, %T\n", k, k)
	}

	line_break()

	// Other variables:
	{
		var n bool = 1 == 1
		fmt.Printf("%v, %T\n", n, n)

		var o float32 = 3.14
		o = 13.7e30
		fmt.Printf("%v, %T\n", o, o)

		var p complex64 = 1 + 2i
		fmt.Printf("%v, %T\n", p, p)

		var q string = "this is a string"
		fmt.Printf("%v, %T\n", string(q[2]), q[2])
		r := []byte(q) // bytestring
		fmt.Printf("%v, %T\n", r, r)
		var s rune = 'a'
		fmt.Printf("%v, %T\n", s, s)

		// bit operations:
		a := 10             // 1010
		b := 3              // 0011
		fmt.Println(a & b)  // AND    (0010)
		fmt.Println(a | b)  // OR     (1011)
		fmt.Println(a ^ b)  // XOR    (1001)
		fmt.Println(a &^ b) // ANDNOT (0100) FLIPS THE SECOND VAR (b) THEN DOES AND

		// bit shifting:
		c := 8              // 1000
		fmt.Println(c << 3) // 1000000 (SHIFTING LEFT = LARGER)
		fmt.Println(c >> 3) // 0001 (SHIFTING RIGHT = SMALLER	)

	}

	line_break()

	{
		fileSize := 4000000000.0
		fmt.Printf("%.2fGB\n", fileSize/GB)
	}

	// VARIABLES HAVE TO BE USED: IF NOT, COMPILER ERROR
	// VARIABLE VERBOSITY DEPEND ON ITS SCOPE (block = small names, package= med names, global = long names)
	// ACRONYMS SHOULD BE ALL UPPERCASE (theURL not theUrl)
}
