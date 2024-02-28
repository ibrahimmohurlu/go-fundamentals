# Strings and Runes

Go treats strings as byte arrays. Strings can be indexed just like an array. Go uses UTF-8 encoding. Strings are immutable in Go so once we define a string we cannot modify it

```go
var myString = "résumé"
var indexed = myString[0]
fmt.Println(indexed) // prints 114 ASCII value of 'r'
for i, v := range myString {
	fmt.Println(i, v) // printing index and value
    // index 2 skipped because é takes 2 bytes to store
    // range keyword takes care of that
}
// prints 8 because 6 characters + 2 extra ones
fmt.Printf("length of 'résumé' is %v\n", len(myString))

var runes = []rune("résumé")
fmt.Println(runes) // [114 233 115 117 109 233] unicodes

// string concatination (not efficient)
var stringArr = []string{"h", "e", "l", "l", "o"}
var str = ""
for i := range stringArr {
	str += stringArr[i]
}
fmt.Println(str) // prints 'hello'

// building string using 'strings' package (more efficient)
var stringBuilder strings.Builder
for i := range stringArr {
	stringBuilder.WriteString(stringArr[i])
}
str = stringBuilder.String()
fmt.Println(str) // prints (hello)
```
