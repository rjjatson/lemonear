package model_test

import (
	"container/heap"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"rjjatson.com/lemonear/internal/model"
)

func TestScoreHeap_SortKLinear(t *testing.T) {
	h := model.NewScoreHeap()

	set := []model.Verdict{
		{
			Score: 7.312,
		},
		{
			Score: 2.6,
		},
		{
			Score: 6.66,
		},
		{
			Score: 3.1,
		},
		{
			Score: 0.3,
		},
	}

	k := 3
	for _, s := range set {
		heap.Push(h, s)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	expectedResult := []float64{3.1, 2.6, 0.3}
	actualResult := []float64{}
	for i := 0; i < k; i++ {
		actualResult = append(actualResult, heap.Pop(h).(model.Verdict).Score)
	}

	assert.Equal(t, expectedResult, actualResult)
}

func TestScoreHeap_SortKParallel(t *testing.T) {
	h := model.NewScoreHeap()

	set := []model.Verdict{
		{
			Score: 7.312,
		},
		{
			Score: 2.6,
		},
		{
			Score: 6.66,
		},
		{
			Score: 3.1,
		},
		{
			Score: 0.3,
		},
	}

	k := 3
	wg := sync.WaitGroup{}
	for _, s := range set {
		wg.Add(1)
		go func(q model.Verdict) {
			h.Lock()
			heap.Push(h, q)
			if h.Len() > k {
				heap.Pop(h)
			}
			h.Unlock()
			wg.Done()
		}(s)
	}

	wg.Wait()

	expectedResult := []float64{3.1, 2.6, 0.3}
	actualResult := []float64{}
	for i := 0; i < k; i++ {
		actualResult = append(actualResult, heap.Pop(h).(model.Verdict).Score)
	}

	assert.Equal(t, expectedResult, actualResult)
}
