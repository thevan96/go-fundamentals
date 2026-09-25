package main

import "fmt"

type Point struct {
	X, Y int
}

func Distance(p1, p2 Point) int {
	deltaX := (p1.X -p2.X)
	deltaY := (p1.Y -p2.Y)
	return deltaX*deltaX + deltaY *deltaY
}

func main() {
	var x1, y1, x2, y2 int
	fmt.Scan(&x1)
	fmt.Scan(&y1)
	fmt.Scan(&x2)
	fmt.Scan(&y2)

	p1 := Point{x1, y1}
	p2 := Point{x2, y2}
	fmt.Println(Distance(p1, p2))
}
