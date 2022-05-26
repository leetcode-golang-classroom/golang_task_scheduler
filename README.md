# golang_task_scheduler

Given a characters array `tasks`, representing the tasks a CPU needs to do, where each letter represents a different task. Tasks could be done in any order. Each task is done in one unit of time. For each unit of time, the CPU could complete either one task or just be idle.

However, there is a non-negative integer `n` that represents the cooldown period between two **same tasks** (the same letter in the array), that is that there must be at least `n` units of time between any two same tasks.

Return *the least number of units of times that the CPU will take to finish all the given tasks*.

## Examples

**Example 1:**

```
Input: tasks = ["A","A","A","B","B","B"], n = 2
Output: 8
Explanation:
A -> B -> idle -> A -> B -> idle -> A -> B
There is at least 2 units of time between any two same tasks.

```

**Example 2:**

```
Input: tasks = ["A","A","A","B","B","B"], n = 0
Output: 6
Explanation: On this case any permutation of size 6 would work since n = 0.
["A","A","A","B","B","B"]
["A","B","A","B","A","B"]
["B","B","B","A","A","A"]
...
And so on.

```

**Example 3:**

```
Input: tasks = ["A","A","A","A","A","A","B","C","D","E","F","G"], n = 2
Output: 16
Explanation:
One possible solution is
A -> B -> C -> A -> D -> E -> A -> F -> G -> A -> idle -> idle -> A -> idle -> idle -> A

```

**Constraints:**

- `1 <= task.length <= $10^4$`
- `tasks[i]` is upper-case English letter.
- The integer `n` is in the range `[0, 100]`.

## 解析

題目給定一個字元陣列 tasks , 每個字元代表一個 task, 還有一個 n 代表每個相同類型 task 需要間隔的單位時間

每個不同字元代表不同的 task

要寫一個演算法去找出最佳的執行時間

題目關鍵點在於相同 tasks 間需要 idle n 個單位時間

所以看出瓶頸會是出現頻率最高的 tasks

策略有以下特點

1. 每次都要以當下還沒執行 tasks 中出現最多的那個先拿出來執行
2. 每次執行後要紀錄該類 tasks 需要 idle n 單位時間

把每個類別計次

然後把次數放入 MaxHeap

這樣每次都可以找到當下最多次數的 tasks 

執行完後放到一個 IdleQueue 並且紀錄該類 tasks 剩餘個數 還有需要 idle 到哪個時間點

這樣每次都先從 MaxHeap pop 出最多次數的 tasks 出來執行

如果執行完還有剩餘 就放到 IdleQueue並紀錄下次要執行的時間點當下累計時間+n

這樣第一次放到 MaxHeap  需要 O(n * log(26)) 

每次 Pop 出來 log(26)

最差的狀況是 O(n *m) 假設都是同一類 Tasks ide 時間為 m

 

## 程式碼

```go
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

```

## 困難點

1. 理解如何處理 idle tasks 這邊採用 queue
2. 需要理解最佳化的策略是先把出現最多次的 tasks 先選出來做處理

## Solve Point

- [x]  Understand what problem need to solve
- [x]  Analysis Complexity