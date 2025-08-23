package task4

import "math"

type Circle struct {
	JariJari float64
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.JariJari, 2)
}
