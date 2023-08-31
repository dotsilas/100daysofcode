package main

import (
	"fmt"
	_ "sync" // not used package
)

// About the language

func packages() {
	// packages == modules or libraries in other languages
	// package main and func main() {} is the entry point of programs

	// use of "fmt" packages
	fmt.Println("Hello world!")
}

func variables() {
	// variable creation and assignation
	var a, b, c int
	a = 10
	b = 20
	c = 30

	// type inference assignation
	aText := "This is a text"
	result := a + b + c

	// string merging placeholders: %s strings; %d decimal; %f flotante; %t boolean; %% percentage simbol
	fmt.Printf("%s and %d + %d + %d = %d\n", aText, a, b, c, result)
}

func dataTypes() {
	var (
		integer    int     = 10
		float      float64 = 10.23
		aString    string  = "a text stored in a variable"
		isVariable bool    = true
	)

	var integerNotDefined int // zero "0" by dafault
	fmt.Printf("Data types: int=%d, float=%f, string=%s, boolean=%t, not-assigned int=%d",
		integer,
		float,
		aString,
		isVariable,
		integerNotDefined,
	)
}

func loops() {
	// go only has for loops

	// classic C-style for
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// making i accesible outside loop
	var i int
	for ; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("i's value is ", i)

	// while-like loop with for var n int
	n := 3
	for n < 5 {
		fmt.Println("n is", n)
		n++
	}

	// infinite loops
	// for {
	// 	fmt.Println("hello!")
	// }

	// use of break and continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

func conditionals() {
	x := 3
	// if statement
	if x > 2 {
		fmt.Println("x is greater than 2")
	}

	// if statement w/ init statement
	// if err := someFunction(); err != nil {
	// 	fmt.Print(err)
	// }

	if true {
		fmt.Println("Is true")
	} else {
		fmt.Println("Isn't true")
	}

	// "else" omition
	// v, err := someFunc()
	// if someFunc != nil {
	// 	return err
	// }
	// fmt.Println(v)
	// return nil

	// else if
	if x > 0 {
		fmt.Println("x is greater than 0")
	} else if x < 0 {
		fmt.Println("x is less than 0")
	} else {
		fmt.Println("x is equal to 0")
	}

	// switch statement
	switch x {
	case 3:
		fmt.Println("x is 3")
	case 4, 5:
		fmt.Println("x is 4 or 5")
	default: // in Go default is optional
		fmt.Println("x is unknown")
	}

	// switch declaration + match value
	// switch x:= someFunction(); x {
	// case "none":
	// 	fmt.Println("yes")
	// case "yes":
	// 	fmt.Println("not yes")
	// }

	switch { // <-- match value not included
	case x > 0:
		fmt.Println("x is greater than 0")
	case x < 0:
		fmt.Println("x is less than 0")
	default:
		fmt.Println("x is 0")
	}
}

// FUNCTIONS
// is not possible to declarate regular functions inside functions
// but is possible to declarate anonimous functions inside other functions
func add(x, y int) int {
	return x + y
}

// functions with multiple return

// same type params = only one type declaration
func divide(num, div int) (res, rem int) {
	// res and rem don't use :=
	res = num / div
	rem = num % div
	// multiple return
	return res, rem
}

// variadic arguments
func sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

// variadic notation
func sum2(numbers ...int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

// Public and Private
var aPrivateVariable = "" // Private: lowerCase start letter
var APublicVariable = ""  // Public: UpperCase start letter

func arrays() {
	// arrays are length fixed

	// 3 integers length array
	anArray := [3]int{1, 2, 3}
	fmt.Println(anArray)

	// accessing/modifiying arrays
	anArray[2] = 5
	fmt.Println(anArray[1])

}

// Arrays passed in function are copies
func changeValueAtZeroIndex(array [2]int) {
	array[0] = 3
	fmt.Println("inside: ", array[0]) // Will print 3
}

func slices() {
	// like arrays but not statically sized
	aSlice := []int{1, 2, 3, 4}

	// accessing/modifiying slices
	aSlice[1] = 3
	fmt.Println(aSlice)

	// add elements to slice
	aSlice = append(aSlice, 15, 16)
	fmt.Println(aSlice)

	// slicing slices:
	// slice[inclusive index : exclusive index]
	aNewSlice := aSlice[2:4]
	fmt.Println(aNewSlice)

	// slice are views: sliced slices change original slices
	aNewSlice = append(aNewSlice, 123)
	fmt.Println(aNewSlice)
	fmt.Println(aSlice)

	// slice iteration
	iterableSlice := []string{"first", "second", "third", "fourth"}

	// get indices and values of slice with range
	// _ for omiting one return of range
	for _, value := range iterableSlice {
		fmt.Println(value)
	}
}

func maps() {
	// maps: key, value structure
	// declarations of maps

	// with "make(map[key type]value type, size)"
	// size can be omited
	aMap := make(map[string]int, 2)
	aMap["number"] = 1
	aMap["age"] = 20
	fmt.Println(aMap)

	// direct assignation
	otherMap := map[string]string{
		"name": "Jesus",
		"age":  "2023",
	}
	fmt.Println(otherMap)
}

func pointers() {
	// get memory address with &
	x := 10
	fmt.Printf("x address = %p\n", &x)

	// declaring a pointer variable
	var xPointer *int
	xPointer = &x
	fmt.Printf("x address = %p\n", xPointer)

	// modifying value
	fmt.Println(x)         // 10
	fmt.Println(*xPointer) // 10 (dereferencing the pointer)
	*xPointer = 5          // Modifying the value via the pointer
	fmt.Println(x)         // 5 (the original variable's value is changed)

	// Note:
	// Functions by dafault make copies of arguments.
	// They don't affect original variables.
}

func changeValue(x string) {
	x += "different text"
}

func myStructs() {
	// structs == collection of variables
	var record = struct {
		name string
		age  int
	}{
		name: "Jesus",
		age:  2023,
	}
	fmt.Printf("Name: %s; Age: %d\n", record.name, record.age)

	// custom types
	type CarModel string
	var myCar CarModel = "Mustang"
	var myOtherCar = CarModel("Toyota") // by type conversion
	fmt.Println(myCar, " ", myOtherCar)

	jesus := Record{Name: "Jesus", Age: 2023}
	// fmt.Printf("Name: %s; Age: %d\n", jesus.Name, jesus.Age)

	fmt.Println(jesus.String())
}

// custom struct types
type Record struct {
	Name string
	Age  int
}

// methods == functions attached to a data type
// behaviour of a struct
func (r Record) String() string {
	return fmt.Sprintf("%s,%d", r.Name, r.Age)
}

// Constructor methods
func NewRecord(name string, age int) (*Record, error) {
	if name == "" {
		fmt.Println("name cannot be empty string")
	}
	if age <= 0 {
		fmt.Println("age cannot be <= 0")
	}
	return &Record{Name: name, Age: age}, nil
}

// // Interfaces
// // stores any value that declares a set of methods
// type Stringer interface {
// 	String() string
// }

// type Person struct {
// 	first, last string
// }

// func (p Person) String() string {
// 	return fmt.Sprintf("%s, %s", p.last, p.first)
// }

// type StrList []string

// func (s StrList) String() string {
// 	return strings.Join(s, ",")
// }

// interface declaration
type Writer interface {
	// abstract method:
	// declared but not implemented
	Write() string
}

// first type
type Novelist struct {
	name, message string
}

// second type
type Programmer struct {
	name, message string
}

// Novelist implementation of Writter interface
func (n Novelist) Write() string {
	return n.message
}


// Programmer implementation of Writter interface
func (p Programmer) Write() string {
	return p.message
}

// Novelist and Programmer implements
// all methods of Writer (Write), so:
// Novelist and Programmer satisfy Writter

func main() {
	// packages()
	// variables()
	// dataTypes()
	// loops()
	// conditionals()

	result := add(10, 20)
	fmt.Printf("10 + 20 = %d\n", result)

	// if you only wants one return, can use _ to not used
	result, remainder := divide(10, 3)
	fmt.Printf("10 / 3 = %d and remainder %d\n", result, remainder)

	sum2Result := sum2(1, 2, 3, 4, 5)
	fmt.Println(sum2Result)

	// anonimous function
	greet := func() string {
		return "Hello!"
	}

	fmt.Println(greet())

	// create a single-use function and pass arguments
	anotherGreet := func(word1, word2 string) string {
		return word1 + " " + word2
	}("hola", "mundo")

	fmt.Println(anotherGreet)

	fmt.Println(aPrivateVariable)

	arrays()

	// using arrays in functions make copies
	arr := [2]int{}
	changeValueAtZeroIndex(arr)
	fmt.Println(arr)

	// slices()
	// maps()
	// pointers()
	// myStructs()
}
