package worker

import (
	"container/heap"

	"rjjatson.com/lemonear/internal/model"
	"rjjatson.com/lemonear/internal/service/worker/algorithm"
	"rjjatson.com/lemonear/internal/service/worker/algorithm/euclidean"
)

type Worker struct {
	maxDistance float64
	maxResult   int
	maxWorker   int
	inChan      chan *model.Job
	stopChan    chan bool
	algorithm   algorithm.Algorithm
}

// set max result and max distance to 0 if no max was set
func New(maxWorker int, maxResult int, maxDistance float64) *Worker {
	return &Worker{
		maxDistance: maxDistance,
		maxResult:   maxResult,
		maxWorker:   maxWorker,
		inChan:      make(chan *model.Job),
		stopChan:    make(chan bool),
		algorithm:   euclidean.NewEuclidean(), // plug and play
	}
}

func (w *Worker) Run() {
	for i := 0; i < w.maxWorker; i++ {
		go w.listen()
	}
}

func (w *Worker) Stop() {
	close(w.stopChan)
}

func (w *Worker) Queue(job *model.Job) {
	w.inChan <- job
}

func (w *Worker) listen() {
	stop := false
	for {
		if stop {
			break
		}
		select {
		case job := <-w.inChan:
			w.proccess(job)
		case stop = <-w.stopChan:
		}
	}
}

func (w *Worker) proccess(job *model.Job) {
	defer job.Done()

	d := w.algorithm.CalculateDistance(job.Origin, job.Destination)
	if w.maxDistance != 0 && d > w.maxDistance {
		return
	}

	v := model.Verdict{
		Score: d,
		Point: job.Destination,
	}

	job.Result.Lock()
	heap.Push(job.Result, v)
	if w.maxResult != 0 && job.Result.Len() > w.maxResult {
		heap.Pop(job.Result)
	}
	job.Result.Unlock()
}
