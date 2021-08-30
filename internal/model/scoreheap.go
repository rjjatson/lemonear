package model

import "sync"

type ScoreHeap struct {
	mu      *sync.Mutex
	verdict []Verdict
}

func NewScoreHeap() *ScoreHeap {
	return &ScoreHeap{
		mu:      &sync.Mutex{},
		verdict: make([]Verdict, 0),
	}
}

func (h *ScoreHeap) Lock() {
	h.mu.Lock()
}

func (h *ScoreHeap) Unlock() {
	h.mu.Unlock()
}

func (h *ScoreHeap) Len() int {
	return len(h.verdict)
}

func (h *ScoreHeap) Less(i, j int) bool {
	r := h.verdict[i].Score > h.verdict[j].Score
	return r
}

func (h *ScoreHeap) Swap(i, j int) {
	h.verdict[i], h.verdict[j] = h.verdict[j], h.verdict[i]
}

func (h *ScoreHeap) Push(x interface{}) {
	h.verdict = append(h.verdict, x.(Verdict))
}

func (h *ScoreHeap) Pop() interface{} {
	old := h.verdict
	n := len(old)
	x := old[n-1]
	h.verdict = old[0 : n-1]
	return x
}
