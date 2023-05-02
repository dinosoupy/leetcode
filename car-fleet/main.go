package main

import (
	"fmt"
	"sort"
)

func main() {
	pos := []int{10,8,0,5,3}
	speed := []int{2,4,1,1,3}
	target := 12
	fmt.Println(carFleet(target, pos, speed))
}

func carFleet(target int, position []int, speed []int) int {
	// Array of cars with (position, speed) as elements
	cars := make([][2]int, len(position))

	// Parse positions and speed into cars
	for i := 0; i < len(position); i++ {
		cars[i] = [2]int{position[i], speed[i]}
	}

	// Sort cars by desc order of position (closest to target first)
	sort.Slice(cars, func(i, j int) bool {
		return cars[i][0] > cars[j][0]
	})

	var currentSlowest float64
	var fleet int

	for i := 0; i < len(position); i++ {
		timeToTarget := float64(target - cars[i][0])/ float64(cars[i][1])
		// Loop through each cars timeToTarget; since timeToTarget should be monotonic non-decreasing for the cars array
		// any violation is counted as a fleet being added
		if timeToTarget > currentSlowest {
			currentSlowest = timeToTarget
			fleet ++
		}
	}
	return fleet
}