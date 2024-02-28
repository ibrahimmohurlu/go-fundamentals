# Variables and Data Types

## Variables

In Go `var <*name> <*type>` defines a variable. There are also other ways to define a variable:

```go
var myVar1 string // gets default value of "" (empty string)
var myVar2 = "hi" // type is inferred as string
myVar3 := 123 // type is inferred as int
returnVal := foo() // can't see the type directly bad for readability
```

## Data Types

Data types specify the type of data that a valid Go variable can hold. In Go language, the type is divided into four categories which are as follows:

1. <b>Basic type:</b> <i>Numbers (default value is 0)</i>, <i>strings (default values is "")</i>, and <i>booleans (default values is false)</i> come under this category.
2. <b>Aggregate type:</b> Array and structs come under this category.
3. <b>Reference type:</b> Pointers, slices, maps, functions, and channels come under this category.
4. <b>Interface type</b>

### Basic Data Types (numbers, strings and booleans)

| Data Type  | Description                                                                                                 |
| ---------- | ----------------------------------------------------------------------------------------------------------- |
| int8       | 8-bit signed integer                                                                                        |
| int16      | 16-bit signed integer                                                                                       |
| int32      | 32-bit signed integer                                                                                       |
| int64      | 64-bit signed integer                                                                                       |
| uint8      | 8-bit unsigned integer                                                                                      |
| uint16     | 16-bit unsigned integer                                                                                     |
| uint32     | 32-bit unsigned integer                                                                                     |
| uint64     | 64-bit unsigned integer                                                                                     |
| int        | Both int and uint contain same size, either 32 or 64 bit (CPU architecture determines the size).            |
| uint       | Both int and uint contain same size, either 32 or 64 bit (CPU architecture determines the size).            |
| rune       | It is a synonym of int32 and also represent Unicode code points.                                            |
| byte       | It is a synonym of uint8.                                                                                   |
| uintptr    | It is an unsigned integer type. Its width is not defined, but its can hold all the bits of a pointer value. |
| float32    | 32-bit IEEE 754 floating-point number                                                                       |
| float64    | 64-bit IEEE 754 floating-point number                                                                       |
| complex64  | Complex numbers which contain float32 as a real and imaginary component.                                    |
| complex128 | Complex numbers which contain float64 as a real and imaginary component.                                    |
| bool       | 1 bit of an information either true or false                                                                |
| string     | String is a sequence of immutable bytes. Once string is created you cannot change it.                       |
