package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Add(z float64) float64 {
	return v.X + v.Y + z
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v)
	fmt.Println(v.Abs())
	fmt.Println(v.Add(1))

}
