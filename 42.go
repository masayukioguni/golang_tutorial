package main

import (
	"fmt"
)

func main() {
	hypot := func(x, y float64) float64 {
		return (x*x + y*y)
	}

	fmt.Println(hypot(3, 4))
}
