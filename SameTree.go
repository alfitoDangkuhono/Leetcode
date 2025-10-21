package main

import "fmt"

func isSameTree(p []int, q []int) bool {

	var pointer_p int = 0
	var pointer_q int = 0
	var result bool = false

	for x := pointer_p; x < len(p); x++ {
		for c := pointer_q; c < len(q); c++ {
			if p[pointer_p] == q[pointer_q] {
				pointer_p++
				pointer_q++
				result = true
			} else {
				result = false
				break
			}
		}
	}
	if len(p)+len(q) == pointer_p+pointer_q {
		return result
	}
	return false
}

func main() {

	fmt.Print(isSameTree([]int{1, 3}, []int{1, 2, 3}))
}
