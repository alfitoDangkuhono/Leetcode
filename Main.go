package main

import (
	"fmt"
	"strconv"
)

func main() {
	// variable for function
	array := [6]int{1, 2, 3, 4, 5, 6}

	// declare another function file
	insertPosition := searchInsert(array[:], 6)
	LengthOfLastWord := LengthOfLastWord("Hello World")

	fmt.Println("Insert Position " + strconv.Itoa(insertPosition))
	fmt.Println("Length Of Last Word " + strconv.Itoa(LengthOfLastWord))
}
