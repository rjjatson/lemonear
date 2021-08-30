package model

type Point struct {
	X  float64
	Y  float64
	ID int64
}

type Verdict struct {
	Score float64
	Point Point
}
