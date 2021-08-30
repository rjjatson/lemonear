package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"rjjatson.com/lemonear/internal/model"
	"rjjatson.com/lemonear/internal/service"
)

func TestService(t *testing.T) {
	svc := service.New(3, 4, 1.5)
	svc.Run()

	for i := 0; i < 100; i++ {
		svc.InsertPoints(
			model.Point{
				ID: int64(i),
				X:  float64(i),
				Y:  float64(i),
			},
		)
	}

	actualResult, _ := svc.GetNearest(model.Point{
		X: 23,
		Y: 22,
	})

	assert.Contains(t, actualResult,
		model.Verdict{
			Score: 1,
			Point: model.Point{
				ID: 22,
				X:  22,
				Y:  22,
			},
		})

	assert.Contains(t, actualResult,
		model.Verdict{
			Score: 1,
			Point: model.Point{
				ID: 23,
				X:  23,
				Y:  23,
			},
		})
}
