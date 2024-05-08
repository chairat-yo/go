package main

import (
	"fmt"
)

func main() {
	//hello world
	fmt.Println("Hello World!")

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
	fmt.Println(example1)

	var example2 int
	//example2 = 9
	fmt.Println(example2)

	// var x, y, z = 'x', "y", `zzzz`
	// fmt.Println(x)

	//Block Var
	// var (
	//   a int = 1
	//   b int = 2
	//   c = "hi hi"
	// )

	//Create & Assign
	test_var := 4
	fmt.Println(test_var)

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
	fmt.Println(a, b)

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
		fmt.Println("Impossible")
	} else {
		fmt.Println("1<2")
	}

	//if with func
	// if justReturnValue() != 2 {
	// 	fmt.Println("Not 2!!!")
	// }

	// //Statement Initialization
	// if i := 5; i < 10 {

	// }

	// if i, err := justReturnValue(); i < 5 {
	// 	fmt.Println(err)
	// }

	//Early Return
	token, err := justReturnValue()
	if err != "" {
		fmt.Println(token)
		return
	}

	//Switch
	x := 3
	switch x {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("default:", x)
	}

	switch result := calculate(4, 5); { //notice semi colon
	case result == 20:
		fmt.Println("Correct!!!")
	default:
		fmt.Println("Wrong!!!!!")
	}

	//Case list
	switch x {
	case 1, 2, 3:
		fmt.Println("1, 2, 3")
	case 10, 11, 12:
		fmt.Println("Other")
	}

	//fallthrough - execute next case even not match!!!
	switch x {
	case 1:
	case 2:
	case 3:
		fmt.Println("Case 3!!!!")
		fallthrough
	case 4:
		fmt.Println("not match executed....")
		fallthrough
	case 5:
		fmt.Println("more execution")
	default:
		fmt.Println("default")
	}

	//Loop
	//for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//while in go
	i := 0
	for i < 10 {
		fmt.Print(i)
		i++
	}

	//for: Infinite
	fmt.Println()
	for {
		if i == 0 {
			break
		}
		fmt.Print(i, ",")
		i--
	}

	fmt.Println()
	//instantiating a structure
	data1 := Sample{"world", 1, 2}
	// data := Sample{
	// 	field: "word",
	// 	a:     1,
	// 	b:     2,
	// }

	fmt.Println(data1)

	data2 := Sample{field1: "xxx"}
	fmt.Println(data2)

	field1Test := data2.field1
	fmt.Println("field11Test.field1:", field1Test)

	data2.field1 = "new xxx"
	fmt.Println("data2 new field1:", data2)

	//Anonymous Structures
	//anonymous/inline inside function
	//useful with library functions or shipping data accross a network
	//define data structures
	var anonymousSample struct {
		stringField string
		int1, int2  int
	}

	anonymousSample.stringField = "stringField"
	anonymousSample.int1 = 1
	anonymousSample.int2 = 2
	fmt.Println("anonymousSample:", anonymousSample)

	shortSample := struct {
		stringField string
		int1, int2  int
	}{
		"string1",
		1, 2,
	}
	fmt.Println("shortSample:", shortSample)

	//Arrays
	//var myArray [3]int
	//myArray := [3]int{7, 8, 9}
	//myArray := [...]int{7, 8, 9} // ... -> find for us
	//myArray := [4]int{7, 8, 9}
	//var myNewArrayTest [3]int
	myNewArrayTest := [...]int{7, 8, 9, 0, 11, 12, 13, 14, 15}
	for i := 0; i < len(myNewArrayTest); i++ {
		fmt.Println("var size array", i)
	}

	myFixedArrayTest := [3]int{1, 2, 3}
	for i := 0; i < len(myFixedArrayTest); i++ {
		fmt.Println("fixed size array", i)
	}

	//31 +Slices
	//work with arrays
	//"View" for array (dynamic)
	//functions can accept slice
	//small amount of mem
	mySlice := []int{1, 2, 3}
	for i := 0; i < len(mySlice); i++ {
		fmt.Println("slice:", mySlice[i])
	}
	//slice[a:b] a --> start(inclusive), b --> end(exclusive)
	//slice[:]
	//slice[1:]
	//slice[:2]
	//used to create arrays that can be extended
	numbers := []int{1, 2, 3}
	numbers = append(numbers, 4, 5, 6)
	fmt.Println("Appended Array:", numbers)

	part1 := []int{1, 2, 3}
	part2 := []int{4, 5, 6}
	combined := append(part1, part2...)
	fmt.Println("Combined Array:", combined)

	//pre allocation
	slice := make([]int, 10)
	fmt.Println("Allocated Slice:", slice)

	//slices to functions
	//func funcName(slice []int) {}

	//multidimensional slices
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"}}

	fmt.Println(board)

	//35 Map
	//unordered
	//extremely high performance
	myMap := make(map[string]int)

	// myMap := map[string]int{
	// 	"item 1": 1,
	// 	"item 2": 2,
	// 	"item 3": 3,
	// }

	fmt.Println("myMap:", myMap)

	//Map Operations
	//Insert
	myMap["favorite number"] = 5
	//Read
	fav := myMap["favorite number"]
	missing := myMap["age"]
	fmt.Println("favorite number:", fav, "missing:", missing)

	//Delete
	delete(myMap, "favorite number")
	fmt.Println("deleted map:", myMap)

	//Check Map
	noname, found := myMap["noname"]
	if !found {
		fmt.Println("noname not found:", noname)
	}

	//Iteration
	for key, value := range myMap {
		fmt.Println("Key:", key, "Value:", value)
	}

	//Map Demo
	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 11
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1
	shoppingList["eggs"] += 1
	fmt.Println("shoppingList:", shoppingList)

	//Pointers
	//Go is "pass by value"
	//var valuePtr *int
	//valuePtr := &value
	value := 10
	var valuePtr *int
	valuePtr = &value
	fmt.Println("valuePtr:", valuePtr)

	//Dereference Demo
	i = 1
	increment(&i)
	fmt.Println("Dereference value:", i)

	//Lesson 43 Receiver Functions
	//modified function signature --> dot notation

	//iota
	//const
	//iota --> auto assign values
	const ( //short-form
		Online = iota
		Offline
		Maintenance
		Retired

	//49 Text formatting
	//fmt
	//Fprint --> stream
	//Sprint --> new string

	//Prinf verbs
	//%v	default
	//%t	true/false
	//%c	Character
	//%X	Hex
	//%U	Unicode
	//%e	Scientific

	//Escape Sequences
	// \\
	// \'
	// \"
	// \n
	// \u or \U
	// \x Raw bytes
	)

	//Skipping Values
	const (
		s0 = iota
		_
		_
		s3
		s4
	)

	const (
		i3 = iota + 3
		i4
		i5
	)
}

//Dereference Pointers
func increment(x *int) {
	*x += 1
}

//Receiver Function
// func (coord *Coordinate) shiftBy(x, y int) {
// 	coord.X += x
// 	coord.Y += y
// }

//Receiver Value

//var myTestArray = [...]int{7, 8, 9}
// for i := 0; i < 10 i++ {
// 	item := myTestArray[i]
// 	fmt.Println(item)
// }

//Structure
//~class
//field in groups -> more efficient
//can associate functionality
type Sample struct {
	field1 string
	a, b   int
}

func calculate(x int, y int) int {
	return x * y
}

//Functions - camelCase
func name(param1, param2 float32) float32 {
	fmt.Print("Func Name")
	return param1 + param2
}

func multiReturn() (int, int, int) {
	return 0, 1, 2
}

func justReturnValue() (int, string) {
	return 1, ""
}
