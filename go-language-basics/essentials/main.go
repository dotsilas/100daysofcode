package main
import (
	"errors"
	"fmt"
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

func main() {
	dividedBy := []int{0,1,2,3}
	for _, div := range dividedBy {
		res, err := Divide(100, div)
		if err != nil {
			fmt.Printf("100 by %d error: %s\n", div, err)
			continue
		}
		fmt.Printf("100 divided by %d = %d\n", div, res)
	}
}
