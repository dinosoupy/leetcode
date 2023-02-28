package main

import (
	"fmt"
)

func main() {
	test_nums := []int{3,2,3}
	test_target := 6
	fmt.Println(twoSum(test_nums, test_target))
}

func twoSum(nums []int, target int) []int {
	// Hashmap m: element in nums is key and it's (index + 1) is the value
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		// Check if a value exists for the required key; if ok: return
		if j, ok := m[target-num]; ok {
			return []int{j, i}
		}
		// Add the current element to the map
		m[num] = i
	}
	// Return empty array if no solution found
	return []int{}
}

// // Brute force approach with complexity O(n^2)
// func twoSum(nums []int, target int) []int {
// 	// Loop through nums twice
// 	for i, a := range nums {
// 		for j, b := range nums {
// 			// Prevent selection of same element twice; check if sum is target
// 			if i != j && a + b == target {
// 				return []int{i, j}
// 			}
// 		}
// 	} 
// 	// Return empty array if no solution found
// 	return []int{}
// }