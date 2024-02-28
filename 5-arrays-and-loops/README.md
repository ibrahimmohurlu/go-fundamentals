# Arrays and Loops

## Arrays

- Fixed length
- Same type
- Indexable
- Contiguous memory

```go
var arr [3]int32 // fixed length(3), same type (int32)
arr[1] = 123 // indexable
arr[1:3] // returns elements index of [1,3) (end is excluded)

arr2 :=[5]int{1,2,3,4,5} // type is inferred
arr2 :=[...]int{1,2,3,4,5} // type and length are inferred
```

## Slices

Slice wrap arrays to give a more general, powerful and convenient interface to sequence of data. [Source](https://go.dev/doc/effective_go)

```go
// if length is not specified then this is a slice
var intSlice []int = []int{9,8,7}
// adds element
intSlice = append(intSlice,6)
len(intSlice) // returns the length
cap(intSlice) // returns the capacity

intSlice2 := []int{5,4}
// concat slices with spread operator
intSlice = append(intSlice,intSlice2...)

// create slice length of 3 and capacity of 8
var intSlice3 []int = make(int[],3,8)
```

## Map

Store data as key-value pairs

```go
// string keys mapped to uint32
var myMap = make(map[string]uint32)

// define variable and set value
var myMap2 = map[string]float64{"John": 0.36}

// access value by key (returns 0.36)
myMap2["John"]

// key doesnt exists,returns default value of the type (0 in this case)
myMap2["Doe"]

// exists boolean which indicates value exists or not
var val,exists=myMap2["Doe"]

//deletes the element
delete(myMap2,"John")
```

## Loops

```go
for key,value := range myMap{
    // loop over map with key value
}

for index,value := range myArray{
    // loop over map with index value
}

for i :=0; i < 10; i++ {
    // classic for loop
}

// infinite loop with break statement
index:=0
for {
    if index > 10{
        break
    }
    i++
}
```
