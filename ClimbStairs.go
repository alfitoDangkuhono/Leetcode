package main

import "fmt"

func climbStairs(n int) int {

	a, b := 0, 1

	for x := 0; x < n; x++ {
		a, b = b, a+b
		// fmt.Println("a : ", a, "b : ", b)
	}
	return b
}

func main() {
	fmt.Print(climbStairs(4))
}
