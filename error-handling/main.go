package main

import "fmt"

type CustomError struct {
	Msg  string
	Code int
}

func (e CustomError) Error() string {
	return fmt.Sprintf("error: %s, code: %d", e.Msg, e.Code)
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, CustomError{Msg: "cannot divide by zero", Code: 500}
	}

	return a / b, nil
}

func main() {
	_, err := Divide(1, 0)
	println(err.Error())
}
