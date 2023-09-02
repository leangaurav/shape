package shape

type Rectange struct {
	L, B int
}

func (r *Rectange) Area() float64 {
	return float64(r.L * r.B)
}
