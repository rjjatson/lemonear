package service_test

import (
	"math/rand"
	"testing"

	"rjjatson.com/lemonear/internal/model"
	"rjjatson.com/lemonear/internal/service"
)

var svc service.Service

func init() {
	maxWorker := 10
	maxResult := 50
	maxDistance := 10.0
	svc := service.New(maxWorker, maxResult, maxDistance)
	for i := 0; i < 5000000; i++ {
		svc.InsertPoints(model.Point{
			ID: int64(i),
			X:  generateRandomPoint(),
			Y:  generateRandomPoint(),
		})
	}
	svc.Run()
}

func BenchmarkService(b *testing.B) {
	for i := 0; i < b.N; i++ {
		svc.GetNearest(model.Point{
			X: 34.594751,
			Y: 72.394124,
		})
	}
}

func generateRandomPoint() float64 {
	return float64(rand.Intn(100)) + rand.Float64()
}
