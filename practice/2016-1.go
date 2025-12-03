package main

import (
	"fmt"
	"strconv"
	"math"
)

var (
	coordinates = [2]int{0, 0}

	visited = map[[2]int]bool {
		coordinates: true,
	}		
)

func main() {
	var (
		direction = 0 // 0,1,2,3 => N,E,S,W
		sequence = "L3, R1, L4, L1, L2, R4, L3, L3, R2, R3, L5, R1, R3, L4, L1, L2, R2, R1, L4, L4, R2, L5, R3, R2, R1, L1, L2, R2, R2, L1, L1, R2, R1, L3, L5, R4, L3, R3, R3, L5, L190, L4, R4, R51, L4, R5, R5, R2, L1, L3, R1, R4, L3, R1, R3, L5, L4, R2, R5, R2, L1, L5, L1, L1, R78, L3, R2, L3, R5, L2, R2, R4, L1, L4, R1, R185, R3, L4, L1, L1, L3, R4, L4, L1, R5, L5, L1, R5, L1, R2, L5, L2, R4, R3, L2, R3, R1, L3, L5, L4, R3, L2, L4, L5, L4, R1, L1, R5, L2, R4, R2, R3, L1, L1, L4, L3, R4, L3, L5, R2, L5, L1, L1, R2, R3, L5, L3, L2, L1, L4, R4, R4, L2, R3, R1, L2, R1, L2, L2, R3, R3, L1, R4, L5, L3, R4, R4, R1, L2, L5, L3, R1, R4, L2, R5, R4, R2, L5, L3, R4, R1, L1, R5, L3, R1, R5, L2, R1, L5, L2, R2, L2, L3, R3, R3, R1"
		// sequence = "R8, R4, R4, R8"
		distanceStr = ""
	)

	for _, char := range(sequence) {
		var cur = string(char)
		switch(cur) {
			case "R":
				direction = direction + 1
				if direction > 3 {direction = 0}
			case "L":
				direction = direction - 1
				if direction < 0 {direction = 3}
			case ",":
				coordinates = move(direction, distanceStr, coordinates)
				distanceStr = ""
			case " ":
				continue
			default:
				distanceStr += cur
		}
	}

	coordinates = move(direction, distanceStr, coordinates)
	fmt.Printf("Easter Bunny HQ Distance: %d blocks\n", calculateDistance(coordinates))
	return
}

func move(direction int, distanceStr string, coordinates [2]int) [2]int {
	var (
		distance int
		err error
	)

	distance, err = strconv.Atoi(distanceStr)
	if err != nil {
		panic(err)
	}

	for i := 0; i < distance; i++ {
		switch(direction) {
			case 0:
				coordinates[1] += 1
			case 1:
				coordinates[0] += 1
			case 2:
				coordinates[1] -= 1
			case 3:
				coordinates[0] -= 1
			default:
				panic(fmt.Sprintf("unexpected distance: %d", direction))
		}

		if visited[coordinates] {
			fmt.Printf(
				"Been here before: %+v\nDistance: %d blocks\n", 
				coordinates,
				calculateDistance(coordinates),
			)
		} else {
			visited[coordinates] = true
		}
	}

	return coordinates
}

func calculateDistance(coordinates [2]int) int {
	return int(math.Abs(float64(coordinates[0]))) + int(math.Abs(float64(coordinates[1])))
}