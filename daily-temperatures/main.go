package main

import (
    "fmt"
)

func main() {
    test := []int{73,74,75,71,69,72,76,73}
    fmt.Println(dailyTemperatures(test))
}

func dailyTemperatures(temperatures []int) []int {
    // Final output array, sized to the temperature array, initialised to 0
    answer := make([]int, len(temperatures)) 
    // Make a slice which will act as a stack; it stores indexes of 
    // unresolved temperatures
    stack := []int{} 

    // To make it O(n) each temperature's warmer counterpart index is 
    // filled in as we traverse the array
    for i, temp := range(temperatures) {
        // We always check the last element in stack because temperatures 
        // which have been resolved are popped
        for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
            // Get index for the temp which is being resolved
            last := stack[len(stack)-1]
            // Find distance in indexes
            answer[last] = i - last
            // Pop from stack
            stack = stack[:len(stack)-1]
        }
        // Add the current index to stack after every iteration
        stack = append(stack, i)
    }
    return answer
}

// // O(n^2) solution; using brute force to find the answer for each temp
// func dailyTemperatures(temperatures []int) []int {
//     answer := make([]int, len(temperatures))
//     for i, temp1 := range(temperatures) {
//         count := 0
//         foundHigher := false
//         for _, temp2 := range(temperatures[i+1:]) {
//             if temp2 > temp1 {
//                 count++
//                 foundHigher = true
//                 break
//             }
//             if temp2 <= temp1 {
//                 count++
//             }
//         }
//         if foundHigher {
//             answer[i] = count
//         } else {
//             answer[i] = 0
//         }
        
//     }
//     return answer
// }