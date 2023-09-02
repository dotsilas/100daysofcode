package testing 

import (
	"fmt"
)

func Greet(name string) string {
	greeting := "Hello " + name
	return greeting
}

func main() {
	myName := Greet("Silas")
	fmt.Println(myName)
}
