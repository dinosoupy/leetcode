package main

import (
	"fmt"
  "strconv" // need this for converting string to int
)

func main() {
	test := []string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}
    fmt.Println(evalRPN(test))
}

func evalRPN(tokens []string) int {
  answer := []int{}

  for _, token := range(tokens) {
    // For every operator encounered - calculate the result of operation, pop the top 2 numbers, append the result
  	if token == "-" {
  		result := answer[len(answer)-2] - answer[len(answer)-1]
      answer = answer[:(len(answer)-2)]
  		answer = append(answer, result)
  	} else if token == "+" {
  		result := answer[len(answer)-2] + answer[len(answer)-1]
      answer = answer[:(len(answer)-2)]
  		answer = append(answer, result)
  	} else if token == "/" {
  		result := answer[len(answer)-2] / answer[len(answer)-1]
      answer = answer[:(len(answer)-2)]
  		answer = append(answer, result)
  	} else if token == "*" {
  		result := answer[len(answer)-2] * answer[len(answer)-1]
      answer = answer[:(len(answer)-2)]
  		answer = append(answer, result)
  	} else {
  		num, _ := strconv.Atoi(token)
      answer = append(answer, num)
  	}
  }
  return answer[0]
}