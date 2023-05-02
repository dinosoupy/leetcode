package main

import (
	"fmt"
	"sort"
)

func main() {
	pos := []int{0,2,4}
	speed := []int{4,2,1}
	target := 100
	fmt.Println(carFleet(target, pos, speed))
}

func carFleet(target int, position []int, speed []int) int {
	// Create variable n for ease of use
	n := len(position)
	// Make a map[position]speed of length n (where n = number of cars)
  m := make(map[int]int, n)
  // Sort speed array
  speedCopy := speed
  sort.Ints(speedCopy)
  fleets := 0 // Start with 0 fleets
  // Populate map with 0
  for k := 0; k < target+1; k++ {
  	m[k] = 0
  }

  // keep looping until speed of the last car to reach target is the slowest speed amongst all cars
  for m[target] != speedCopy[0] {
  	// Iterate through all n cars
  	for i := 0; i < n; i++ {
  		// Increment position by speed, this is one cycle
  		position[i] += speed[i]
  		// The map will store the least possible speed for a particular position 
  		// If we encounter a speed greater than the speed stored for a position, that means a faster car has encountered
  		// a slower one in its same position (bumper to bumper) and has therefore coalesced into a fleet
  		// At that point we can change the speed of that position to the lowest speed of the two
  		currentSpeed := m[position[i]]
  		// My speed greater than car already in this position
			if speed[i] > currentSpeed {
				speed[i] = currentSpeed
  			m[position[i]] = speed[i]
  			// Increase fleet count by 1
   			fleets += 1
  			
  		} else {
  			// Increment positions for the cars that didnt coalesce into fleets
				m[position[i]] = speed[i]
				position[i] += speed[i]
  		}
		} 
		m[position[i]] = speed[i]
  }
  return fleets
}