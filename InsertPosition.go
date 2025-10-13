package main

func searchInsert(nums []int, target int) int {

	/* --start algorithm--
	   - looping, using for scanning and checking value array nums
	   - then seeting value position if value target not found in array nums and return result position value with compare value target with value nums array
	--end-algorithm--
	*/
	var result int
	var isNone bool = true

	// scanning value array and checking value target
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			result = i
			isNone = false
		}
	}
	// checking value target when not found in array and setting position value with compare value target with value array
	if isNone == true {
		for i := 0; i < len(nums); i++ {
			// compare and setting position value with value target
			if target < nums[i] {
				result = i
				break
			} else if target > nums[i] {
				result = i + 1
			}
		}
	}
	return result
}
