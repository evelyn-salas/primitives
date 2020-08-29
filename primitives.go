package main

import (
	"fmt"
)

func main() {
	/////////////////
	//boolean types//
	/////////////////

	//creates variable of type boolean
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)

	//logic tests using booleans
	x := 1 == 1
	y := 1 == 2
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)

	//note that uninitialized booleans are set to
	//false by default
	//for example:
	//var n bool
	//n is = 0, aka false

	/////////////////
	//numeric types//
	/////////////////

	//integers
	m := 42 //creates an int by default (if whole)
	fmt.Printf("%v, %T\n", m, m)
	//types of ints: int, int8, int16, int32 and int 64
	//the bigger the number the bigger the range
	//if you need bigger than int64, intstall math package

	//unsigned integers
	//non-zero integers can be put into this type
	//uint8, uint16
	//basically the same as int

	//byte
	//an unsigned 8-bit integer, has it's own type
	//because it is used a lot in programming and
	//it prevents mix ups with ints and uints which
	//may be used in maths

	//basic opperations
	a := 10 //1010
	b := 3  //0011
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b) //so this will give u 3, no decimals
	fmt.Println(a % b) //we get mod in go!

	//different types
	//var i int = 10
	//var k int8 = 3
	//fmt.Println(i + k) //gives an error
	//type conversion is needed (see hello.go)

	//Logical opperators, pertain to binary
	fmt.Println(a & b)  // 0010 = 2
	fmt.Println(a | b)  // 1011 = 11
	fmt.Println(a ^ b)  // 1001 = 9 (exclusivie or)
	fmt.Println(a &^ b) // 0100 = 8 (if niether has it)
	//need to look more into this if it becomes relevant
	//not fully understanding it

	//bit shifting
	t := 8              // = 2^3
	fmt.Println(t << 3) // = 2^3 * 2^3 = 2^6 = 64
	fmt.Println(t >> 3) // = 2^3 / 2^3 = 2^0 = 1

	//floating point numbers
	//float32, float62 (more storage w/ 64)
	p := 3.14   //initialized to float64 by defalut
	p = 13.7e72 //we can use scientific notation
	p = 2.1e14
	fmt.Printf("%v, %T\n", p, p)
	//%, >>, and << only work on int types

	//complex numbers
	var c complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", c, c)
	//we have 2 complex types 64 and 128
	//same arithmetic equations apply
	//to deconstruct
	fmt.Printf("%v, %T\n", real(c), real(c)) //pulls out real part in float32
	fmt.Printf("%v, %T\n", imag(c), imag(c)) //pulls out imag part in float32
	//complex64 -> float32
	//complex128 -> float64
	//compex function = complex(real, imag)

	//////////////
	//text types//
	//////////////

	//Strings
	s := "I am a string"
	fmt.Printf("%v, %T\n", s, s)
	//you can treat this as an array
	fmt.Printf("%v, %T\n", string(s[2]), s[2]) //reads as byte, use string
	//string concatination						//function to get letter
	s2 := " I'm a string too!!"
	fmt.Printf("%v, %T\n", s+s2, s+s2)
	//convert strings to a slice of bytes
	q := []byte(s)
	fmt.Printf("%v, %T\n", q, q)
	//we use a lot of byte slices in go as opposed to strings

	//runes
	//string = untf8 character, rune = untf32 character
	//all stings can be runes, so under 32bytes
	//makes it a little tricky to understand
	r := 'a' //single quotes for a rune, double for string
	fmt.Printf("%v, %T\n", r, r)

}
