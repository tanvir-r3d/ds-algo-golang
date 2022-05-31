package main

import "fmt"

func filter(nums []int, callback func(value int) bool) []int {
	var filtered []int
	for _, value := range nums {
		if callback(value) {
			filtered = append(filtered, value)
		}
	}
	return filtered
}

func main() {
	nums := []int{4, 5, 7, 8, 10}
	nums = filter(nums, func(num int) bool {
		return num%2 == 0
	})

	fmt.Println(nums)
}
