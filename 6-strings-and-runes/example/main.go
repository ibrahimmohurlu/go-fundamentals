package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé"
	var indexed = myString[0]
	fmt.Println(indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Printf("length of 'résumé' is %v\n", len(myString))
	var runes = []rune("résumé")
	fmt.Println(runes)

	// not efficient
	var stringArr = []string{"h", "e", "l", "l", "o"}
	var str = ""
	for i := range stringArr {
		str += stringArr[i]
	}
	fmt.Println(str)

	// efficient
	var stringBuilder strings.Builder
	for i := range stringArr {
		stringBuilder.WriteString(stringArr[i])
	}
	str = stringBuilder.String()
	fmt.Println(str)
}
