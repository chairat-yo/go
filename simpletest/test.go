package main

import (
	halooo "fmt"
)

func main() {
	//hello world
	halooo.Println("Hello World!")

	//data types - statically typed language
	//Go uses type inference
	//Signed Integer: int8, int16, int, int32, int64

	//Unsigned Integer Types: uint8, byte, uint16, uint, uint32, uint64, uintptr

	//Other Types: float32, float64, complex64, complex128, bool

	//Type Aliases:
	//type Speed float64
	//type Velocity Speed

	//Type Conversions
	//type UserId int
	//type Speed float64
	//UserId(7)
	//Speed(99.9)

	//String and Runes
	//Rune - alias for int32
	//Text is represented by rune ~ char
	//String - multiple runes

	//Runes: 'x' 'Y' '3' '\n' '@' `₿`

	//Strings: "Hello\n" "x"

	//Raw Literal: `My name is Chairat "Yo"\n` <-- no new line!!!

	//Go CLI Tool: Update dependencies, Build&test, Manage artifacts, Format source code
	//build-race - check concurrency problems
	//run - run without executable file
	//mod tidy - update dependencies
	//test - run test
	//fmt formats all src

	//Variables
	var example1 = 3
	halooo.Println(example1)

	var example2 int
	//example2 = 9
	halooo.Println(example2)

	// var x, y, z = 'x', "y", `zzzz`
	// halooo.Println(x)

	//Block Var
	// var (
	//   a int = 1
	//   b int = 2
	//   c = "hi hi"
	// )

	//Create & Assign
	test_var := 4
	halooo.Println(test_var)

	//string ("" default)
	//number (0 default)
	//others nil

	//Comma ok (reuse second var)
	//x, err := ...
	//y, err := ...

	//Naming conventions
	//-camelCase
	//-descriptive

	//Constants
	//const MaxValue = 100
	//const Surname = "Nakamoto"
	a, b, _ := multiReturn()
	halooo.Println(a, b)

	//Operators
	// + +=
	// - -=
	// * *=
	// / /=
	// % %=
	// ++ --

	//Logic
	// &&, ||, !

	//if-else
	if 1 > 2 {
		halooo.Println("Impossible")
	} else {
		halooo.Println("1<2")
	}

	//if with func
	if justReturnValue() != 2 {
		halooo.Println("Not 2!!!")
	}

	//Statement Initialization
	if i := 5; i < 10 {

	}

	if i, err := justReturnValue(); i < 5 {
		halooo.Println(err)
	}

	//Early Return
	token, err := justReturnValue()
	if err != "" {
		halooo.Println(token)
		return
	}

	//Switch
	x := 3
	switch x {
	case 1:
		halooo.Println("1")
	case 2:
		halooo.Println("2")
	case 3:
		halooo.Println("3")
	default:
		halooo.Println("default:", x)
	}
}

//Functions - camelCase
func name(param1, param2 float32) float32 {
	halooo.Print("Func Name")
	return param1 + param2
}

func multiReturn() (int, int, int) {
	return 0, 1, 2
}

func justReturnValue() (int, string) {
	return 1, ""
}
