package main

import "fmt"

func main() {
	nums := []int{4, 3, 1, 7}
	target := 10
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		num, found := hashMap[nums[i]]
		if found {
			return []int{num, i}
		}
		hashMap[target-nums[i]] = i
	}

	return []int{}
}
