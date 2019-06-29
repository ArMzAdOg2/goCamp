package main

import (
	"fmt"
	"math"
)

func main() {
	const num int = 10
	open := 0
	var door [num]bool
	for i := 1; i < num; i++ {
		for j := i; j < num; j++ {
			fmt.Println(j, i, math.Mod(float64(j), float64(i)), door[j])
			if math.Mod(float64(j), float64(i)) == 0 {
				if door[j] {
					open--
				} else {
					open++
				}
				door[j] = !door[j]
			}
			fmt.Println(open, door[j])
		}
	}
	fmt.Println(open)
}
