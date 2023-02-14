package main

import (
	"fmt"
)

func main() {
	test_nums := []int{2,3,5,6,6}
	test_target := 12
	fmt.Println(twoSum(test_nums, test_target))
}

func twoSum(nums []int, target int) []int {
	for i, a := range nums {
		for j, b := range nums {
			if i != j && a + b == target {
				return []int{i, j}
			}
		}
	} 
	return []int{}
}