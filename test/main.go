package main

import (
	"area"
	"fmt"
)

func main() {
	a := area.AreaOfCircle(5)

	fmt.Println("Area of circle with radius 5 is ", a)

	fmt.Println("Value of pie is ", area.Pie)

	rect := area.Rectangle{
		Length:  4,
		Breadth: 3,
	}

	aor := area.AreaOfRectangle(rect)

	fmt.Println("Area of rectangle is ", aor)
}
