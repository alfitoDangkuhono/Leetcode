package main

import "fmt"

func TwoSum(nums []int, target int) []int {
	for x := 0; x < len(nums); x++ {
		for c := x + 1; c < len(nums); c++ {
			if nums[x]+nums[c] == target {
				return []int{x, c}
			}
		}
	}
	return []int{}
}

func main() {

	fmt.Print(TwoSum([]int{2, 7, 11, 15}, 9))
}
