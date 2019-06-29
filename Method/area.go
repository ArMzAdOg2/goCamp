package main

import "fmt"

type rectagle struct {
	width  int
	height int
}

type Areaer interface {
	area() float64
}

func (rect rectagle) area() float64 {
	return float64(rect.width) * float64(rect.height)
}

func main() {

	rect := rectagle{3, 4}
	printArea(rect)
}

func printArea(space Areaer) {
	fmt.Println(space.area())
}
