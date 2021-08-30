package euclidean_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"rjjatson.com/lemonear/internal/model"
	"rjjatson.com/lemonear/internal/service/worker/algorithm/euclidean"
)

func TestEuclidean_CalculateDistance(t *testing.T) {
	e := euclidean.NewEuclidean()
	actualResult := e.CalculateDistance(
		model.Point{
			X: 12.32,
			Y: 34.21,
		},
		model.Point{
			X: 0.92,
			Y: 26.55,
		})
	expectedResult := 13.734467590700413
	assert.Equal(t, expectedResult, actualResult)
}
