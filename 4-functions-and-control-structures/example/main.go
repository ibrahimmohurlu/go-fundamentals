package main

import (
	"errors"
	"fmt"
)

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error         // default value is nil (null)
	if denominator == 0 { // trying to divide with 0
		// setting the err (think as throwing error)
		err = errors.New("cant divide with 0")
		// returning
		return 0, 0, err
	}
	var result int = numerator / denominator
	var reminder int = numerator % denominator
	return result, reminder, err
}

func main() {
	var result, reminder, error = intDivision(11, 0)
	// error check
	if error != nil {
		// prints the error message
		fmt.Println(error.Error())
	}
	fmt.Println(result)
	fmt.Println(reminder)
}
