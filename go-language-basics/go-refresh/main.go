package main

import (
	"fmt"
)

/* Parts
1. variables
2. data types
3. constants
4. functions
5. errors
6. if statement
7. switch statement
8. loops
9. slices
10. maps
11. pointers
12. structs
13. custom structs
14. methods
15. constructor method
16. interfaces
17. goroutines
18. testing
19. io
*/

func main() {
	// variable declaration
	var a int

	// variable definition
	a = 1

	// variable declaration and definition
	var b, c int = 2, 3

	// type inference
	subject := "world"

	// data types
	var (
		aInteger int     = 13
		aFloat   float64 = 3.1416
		aString  string  = "hello"
		aBoolean bool    = true
	)

	// constants
	// typed
	const aNumberConstant int = 3
	// untyped: can be used as types of the same family...
	const aStringConstant = "Text Constant"

	// if statements
	if a := true; a {
		fmt.Println("If statement condition is 'true'")
	} else {
		fmt.Println("If statement condition is 'false'")
	}

	// switch
	switch a := 1; a {
	case 1:
		fmt.Println("a is 1")
	case 2, 3:
		fmt.Println("a is 2 or 3")
	default:
		fmt.Println("a is not 1, 2 or 3")
	}

	// switch without match value allows logical valuation
	a = 2
	switch {
	case a > 1:
		fmt.Println("a is greater than 1")
	case a < 1:
		fmt.Println("a is lower than 1")
	default:
		fmt.Println("a isn't greater than 1 or lower than 1")
	}

	fmt.Printf("Hello %s! %d %d %d\n", subject, a, b, c)
	fmt.Printf("Data types: integer %d\nfloat %f\nstring %s\nboolean %t\n", aInteger, aFloat, aString, aBoolean)

	suma := calc(10, 2, "+")
	resta := calc(10, 2, "-")
	mult := calc(10, 2, "*")
	div := calc(10, 2, "/")
	fmt.Printf("Suma: %d, Resta: %d, Multiplicación: %d, División: %d\n", suma, resta, mult, div)

	word := "Otorrinaringologo"
	first, second := divideWord(word)
	fmt.Printf("word: %s, start: %s, tail: %s\n", word, first, second)

	summation := sum(1,2,3,4,5,6,7,8,9)
	fmt.Printf("Summatory of '1,2,3,4,5,6,7,8,9' = %d\n", summation)

	// slices
	sliceOfNames := []string{"John", "Lopez", "Joseph"}
	for _, name := range sliceOfNames {
		fmt.Println(name)
	}
	fmt.Println("===============")

	// slice adding
	sliceOfNames = append(sliceOfNames, "Megan", "Ron", "Harry")
	for _, name := range sliceOfNames {
		fmt.Println(name)
	}
	fmt.Println("===============")

	// slice eliminate element
	sliceOfNames = append(sliceOfNames[:3], sliceOfNames[3+1:]...)
	for _, name := range sliceOfNames {
		fmt.Println(name)
	}
	fmt.Println("===============")

	// maps
	myMap := map[string]string {
		"name":"Jesus",
		"age":"2023",
	}
	fmt.Printf("Name: %s, Age: %s\n", myMap["name"], myMap["age"])

	// pointers
	// structs
	// custom structs
	// methods
	// constructor method
	// interfaces
	// goroutines
	// testing
	// io
}

// public and private data 
var aPrivateData = "" // starts with lowercase
var APrivateData = "" // starts with uppercase 

// functions
func calc(firstNumber, secondNumber int, operator string) int {
	// anonimous function can be used into regular named function, named functions not
	result := func(op string) int {
		switch op {
		case "+":
			return firstNumber + secondNumber
		case "-":
			return firstNumber - secondNumber
		case "*":
			return firstNumber * secondNumber
		case "/":
			return firstNumber / secondNumber
		default:
			fmt.Println("Not valid operator")
			return 0
		}
	}(operator)
	return result
}

// multiple return functions
func divideWord(word string) (start, tail string) {
	length := len(word)
	start = word[:length/2]
	tail = word[length/2:]
	return start, tail
}

// function with variadic arguments
func sum(numbers ...int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}
