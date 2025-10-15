package main

import "fmt"

func removeDuplicates(nums []int) int {

	var result int = 1

	for i := 1; i < len(nums); i++ {
		var isSame bool = false

		for x := 0; x < i; x++ {

			if nums[i] == nums[x] {
				isSame = true
				break
			}
		}
		if !isSame {
			nums[result] = nums[i]
			result++
		}
	}
	return result
}

func main() {

	result := removeDuplicates([]int{0, 0, 0, 1, 2, 3, 1})

	fmt.Println(result)
}
