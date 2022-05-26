package sol

import "container/heap"

type IntMaxHeap []int

// Len() int
func (h *IntMaxHeap) Len() int {
	return len(*h)
}

// Swap(i, j int)
func (h *IntMaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// Less(i, j int) bool
func (h *IntMaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

// Push(val interface{})
func (h *IntMaxHeap) Push(val interface{}) {
	*h = append(*h, val.(int))
}

// Pop() interface{}
func (h *IntMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type IdleTask struct {
	freq     int
	idleTime int
}

func leastInterval(tasks []byte, n int) int {
	// freqMap
	freqMap := make(map[byte]int)
	for _, task := range tasks {
		freqMap[task] += 1
	}
	priorityQueue := &IntMaxHeap{}
	heap.Init(priorityQueue)
	// push task freq into max heap
	for _, freq := range freqMap {
		heap.Push(priorityQueue, freq)
	}
	consumeTime := 0
	idleQueue := []IdleTask{}
	for priorityQueue.Len() > 0 || len(idleQueue) > 0 {
		consumeTime++                // process
		if priorityQueue.Len() > 0 { // process queue
			freq := heap.Pop(priorityQueue).(int)
			if freq-1 > 0 { // still has tasks
				idleQueue = append(idleQueue, IdleTask{freq: freq - 1, idleTime: consumeTime + n})
			}
		}
		if len(idleQueue) > 0 && idleQueue[0].idleTime == consumeTime { // idle task need to process
			heap.Push(priorityQueue, idleQueue[0].freq)
			idleQueue = idleQueue[1:]
		}
	}

	return consumeTime
}
