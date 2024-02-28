# Functions and Control Structures

## Control Structures

Go has <code>if else</code> and <code>switch</code>statement for controlling the flow of the program.

```go
if condition1{
    // this block will be executed if
    // condition is true
}else if condition2{
    // this block will be executed if
    // condition1 is not true and condition2 true
}else{
    //this block will be executed if
    // all conditions above is false
}

switch{
    case condition1 == true && condition2 == false: // equality, and operators
        // this block will be executed if condition is true
    case condition1 != true || condition2 == false: // inequality, or operators
        // this block will be executed if condition is true
    default:
        // thiw block will be executed if none of above is true
}
```

## Functions

In Go we can define a function by <code>func</code> keyword. Functions in Go can have multiple return values. For example we can return errors

```go
func intDivision(numerator int, denominator int) (int, int, error) {
	var err error // default value is nil (null)
	if denominator == 0 { // trying to divide with 0
        // setting the err (think as throwing error)
		err = errors.New("Cant divide with 0")
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
```
