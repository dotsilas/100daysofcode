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
		integer int = 10
		float float64 = 10.23
		aString string = "a text stored in a variable"
		isVariable bool = true
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

	// making variable accesible outside loop
	var i int
	for ;i < 5; i++ {
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
		if i % 2 == 0 {
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

// switch with init statement
	// switch x:= someFunction(); x {
	// case "none":
	// 	fmt.Println("yes")
	// case "yes":
	// 	fmt.Println("not yes")
	// }
	
	switch {
	case x > 0: // boolean evaluations only works if don't use match after switch
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

// multiple return must be placed between ()
func divide(num, div int) (res, rem int) {
	// res and rem don't use := because they are always created
	res = num / div
	rem = num % div
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
aPrivateVariable := "hola" // lowerCase start letter
APublicVariable := "" // UpperCase start letter

func main() {
	packages()
	variables()
	dataTypes()
	loops()
	conditionals()

	result := add(10,20)
	fmt.Printf("10 + 20 = %d\n", result)

	// if you only wants one return, can use _ to not used
	result, remainder := divide(10,3)
	fmt.Printf("10 / 3 = %d and remainder %d\n", result, remainder)

	sum2Result := sum2(1,2,3,4,5)
	fmt.Println(sum2Result)

	// anonimous function
	greet := func() string {
		return "Hello!"
	}

	fmt.Println(greet())

	// create a single-use function and pass arguments
	anotherGreet := func(word1, word2 string) string {
		return word1 + " " + word2
	} ("hola", "mundo")

	fmt.Println(anotherGreet)
}
