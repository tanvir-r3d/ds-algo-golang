package main

import (
	"fmt"
)

func main() {
	nums := []int{-7, -3, 2, 3, 11}

	fmt.Println(squareSortedArray(nums))
}

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func squareSortedArray(nums []int) []int {
	numLen := len(nums)
	resultArray := make([]int, numLen)
	left := 0
	right := numLen - 1
	for i := right; i >= 0; i-- {
		if abs(nums[right]) > abs(nums[left]) {
			resultArray[i] = nums[right] * nums[right]
			right -= 1
		} else {
			resultArray[i] = nums[left] * nums[left]
			left += 1
		}
	}

	return resultArray
}
