package rep

import "math"

type Point struct {
	x, y float64
}

func New() *Point {
	return &Point{}
}

func (p *Point) SetXY(x, y float64) {
	p.x = x
	p.y = y
}

func (p *Point) GetXY() (float64, float64) {
	return p.x, p.y
}
func (p *Point) CountDestination() float64{
	return math.Sqrt(math.Pow(p.x, 2) - math.Pow(p.y, 2))
}
