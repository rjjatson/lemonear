package service

import (
	"container/heap"
	"sync"

	"rjjatson.com/lemonear/internal/model"
	"rjjatson.com/lemonear/internal/service/worker"
)

type Service struct {
	w      *worker.Worker
	points []model.Point
}

func New(maxWorker int, maxResult int, maxDistance float64) *Service {
	return &Service{
		w: worker.New(maxWorker, maxResult, maxDistance),
	}
}

func (s *Service) Run() {
	s.w.Run()
}

func (s *Service) Stop() {
	s.ResetPoints()
	s.w.Stop()
}

func (s *Service) InsertPoints(p model.Point) {
	s.points = append(s.points, p)
}

func (s *Service) ResetPoints() {
	s.points = []model.Point{}
}

func (s *Service) GetNearest(ref model.Point) (nearestPoints []model.Verdict, err error) {
	ver := model.NewScoreHeap()
	wg := &sync.WaitGroup{}
	for _, dst := range s.points {
		wg.Add(1)
		go func(w *sync.WaitGroup, d model.Point) {
			job := model.NewJob(ref, d, ver, w)
			s.w.Queue(job)
		}(wg, dst)
	}

	wg.Wait()
	l := ver.Len()
	for i := 0; i < l; i++ {
		v := heap.Pop(ver).(model.Verdict)
		nearestPoints = append(nearestPoints, v)
	}

	return nearestPoints, nil
}
