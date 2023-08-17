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

	// while-like loop with for
	var n int
	for n < 5 {
		fmt.Println("n is", n)
		n++
	}

	// infinite loops
	for {
		fmt.Println("hello!")
	}
	// )
}

func main() {
	packages()
	variables()
	dataTypes()
	loops()
}
