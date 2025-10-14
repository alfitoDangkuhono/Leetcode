package main

import "fmt"

func plusOne(digits []int) []int {

	for x := len(digits) - 1; x >= 0; x-- {
		if digits[x] < 9 {
			digits[x] += 1
			return digits
		}
		digits[x] = 0
	}
	return append([]int{1}, digits...)
}
func main() {

	fmt.Println(plusOne([]int{9}))
}
