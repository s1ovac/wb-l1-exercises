package point

import "math"

type Point struct {
	x, y float64
}

func (p *Point) SetCoordinates(x, y float64) {
	p.x = x
	p.y = y
}

func (p *Point) Coordinates() (float64, float64) {
	return p.x, p.y
}
func CountDestination(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}
