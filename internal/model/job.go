package model

import "sync"

type Job struct {
	Origin      Point
	Destination Point
	Result      *ScoreHeap
	wg          *sync.WaitGroup
}

func NewJob(
	o Point,
	d Point,
	r *ScoreHeap,
	wg *sync.WaitGroup,
) *Job {
	return &Job{
		Origin:      o,
		Destination: d,
		Result:      r,
		wg:          wg,
	}
}

func (j *Job) Done() {
	j.wg.Done()
}
