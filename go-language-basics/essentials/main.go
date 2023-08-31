package main
import (
	"errors"
	"fmt"
	"log"
)

// things that make Go different:

// error type is an interface
// type error interface {
// 	Error() string
// }

// division by zero error
func Divide(num, div int) (int, error) {
	if div == 0 {
		return 0, errors.New("Cannot divide by 0")
	}
	return num / div, nil
}

// creating named errors
var (
	ErrInput = errors.New("input error")
)

// custom errors
const (
	UnknowCode = 0
	UnreachableCode = 1
	AuthFailureCode = 2
)

type ErrNetwork struct {
	Code int
	Msg string
}

func (e ErrNetwork) Error() string {
	return fmt.Sprintf("network error(%d): %s", e.Code, e.Msg)
}

// constants declaration
// constants can be typed and untyped
const aNumberConstant int64 = 10
const theString = "A string"

// constants as enumaration using 'iota' keyword
const (
	// iota automatically increase number for any new constant
	a = iota // 0
	b // 1
	c // 2
)

// can be used with math operations
const (
	d = iota*2 // 0
	e // 2
	f // 4
)

// untyped const can be used as types of the same family, vars not
type mySpecialStringtype string
func SpecialPrint(s mySpecialStringtype) {
	fmt.Println(string(s))
}

func main() {
	const constString = "hello"
	// var varString string = "hello"

	SpecialPrint(constString)
	// SpecialPrint(varString) // error string != mySpecialStringtype

	dividedBy := []int{0,1,2,3}
	for _, div := range dividedBy {
		res, err := Divide(100, div)
		if err != nil {
			fmt.Printf("100 by %d error: %s\n", div, err)
			continue
		}
		fmt.Printf("100 divided by %d = %d\n", div, res)
	}

	// detect an error with errors.As()
	// var netErr ErrNetwork
	// if errors.As(err, &netErr) {
	// 	if netErr.Code == AuthFailureCode {
	// 	log.Println("unrecoverable auth failure: ", err)
	// 	break
	// 	}
	// 	log.Println("recoverable error: %s", netErr)
	// }
	// log.Println("unrecoverable error: %s", err)
	// break
}

// defer
// executes at the end of a function's execution process
func printStuff() (value string) {
	defer fmt.Println("exiting")
	defer func() {
		value = "we returned this"
	}()
	fmt.Println("I am printing stuff")
	return ""
}

// panic
// stops programm execution
// panic("some bug ocurrence")

// recover
// recover from a panic
func someFunc() {
	defer func() {
		// 
		if r := recover(); r != nil {
			log.Printf("called recover, panic was: %q", r)
		}
	}()
	panic("oh no!!!")
}
