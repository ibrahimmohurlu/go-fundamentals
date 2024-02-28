# Main Points About Go

- Statically Typed Language ( define variables with its types )

  ```go
  var myVar int = 26
  ```

- Strongly Typed Language ( int + float is not allowed)
  ```go
  var n1 int = 36
  var n2 float64 = 0.36
  fmt.Println(n1 + n2) //throws type mismatch error
  ```
- Go is compiled language (compiled the program to executable binary)
- Fast compile time
- Built in concurrency ( Goroutines )
- Designed for simplicity
- Has garbage collector
- Unused imports and variables are not allowed ( throws error when compiling )
