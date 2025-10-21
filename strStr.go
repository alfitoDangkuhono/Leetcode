package main

import "fmt"

func strStr(haystack string, needle string) int {

	var resultPointer int = 0

	for x := 0; x < len(haystack); x++ {
		if haystack[x] == needle {
			resultPointer++
			return resultPointer
		}
	}
	return -1
}

func main() {

	fmt.Print("str : ", strStr("sadbutsad", "sad"))
}
