package euclidean

import (
	"math"

	"rjjatson.com/lemonear/internal/model"
)

type Euclidean struct {
}

func NewEuclidean() *Euclidean {
	return &Euclidean{}
}

func (e *Euclidean) CalculateDistance(src model.Point, dst model.Point) float64 {
	return math.Sqrt(math.Pow((dst.X-src.X), 2) + math.Pow((dst.Y-src.Y), 2))
}
