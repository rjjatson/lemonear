package algorithm

import "rjjatson.com/lemonear/internal/model"

type Algorithm interface {
	CalculateDistance(src model.Point, dst model.Point) float64
}
