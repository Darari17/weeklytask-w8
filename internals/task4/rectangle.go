package task4

type Rectangle struct {
	Panjang float64
	Lebar   float64
}

func (r *Rectangle) Area() float64 {
	return r.Panjang * r.Lebar
}
