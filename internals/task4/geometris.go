package task4

import "fmt"

type Shape interface {
	Area() float64
}

func Kalkulator(shapes []Shape) (result float64) {
	for _, shape := range shapes {
		result += shape.Area()
	}
	return result
}

func DisplayGeometris() {
	circle := Circle{JariJari: 10}
	rectangle := Rectangle{Lebar: 10, Panjang: 20}
	fmt.Printf("Total area: %.2f\n", Kalkulator([]Shape{&circle, &rectangle}))
}
