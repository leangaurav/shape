package shape

import "fmt"

type Circle struct {
	R int
}

func (c *Circle) Area() float64 {
	return 3.14 * float64(c.R*c.R)
}

func NewCircle(r int) (Circle, error) {
	if r <= 0 {
		return Circle{}, fmt.Errorf("radius cannot be <= 0")
	}
	return Circle{r}, nil
}
