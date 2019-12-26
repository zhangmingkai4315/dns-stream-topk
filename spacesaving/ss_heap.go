package spacesaving

// import "container/heap"

type StreamHeap struct {
	capacity      int
	topK          int
	streamCounter uint64
	cache         map[string]*Counter
	heap          *Heap
}

func NewStreamHeap(topk int) *StreamHeap {
	return &StreamHeap{
		capacity:      topk * 2,
		topK:          topk,
		streamCounter: 0,
		cache:         make(map[string]*Counter),
		heap:          NewHeap(topk * 2),
	}
}

func (sh *StreamHeap) existAndUpdate(key string, increment uint64) int {
	for index, item := range sh.heap.Data {
		if item.Key == key {
			item.Counter += increment
			return index
		}
	}
	return -1
}
func (sh *StreamHeap) Name() string {
	return "StreamHeap"
}
func (sh *StreamHeap) Offer(key string, increment int) {
	sh.streamCounter++
	if sh.heap.IsEmpty() {
		sh.heap.Push(NewHeapNode(key, uint64(increment)))
	}
	if index := sh.existAndUpdate(key, uint64(increment)); index != -1 {
		sh.heap.Fix(index)
		return
	}
	if sh.heap.IsFull() {
		min, _ := sh.heap.Min()
		min.Key = key
		min.Error = min.Counter
		min.Counter = min.Counter + uint64(increment)
		sh.heap.Fix(0)
	} else {
		sh.heap.Push(NewHeapNode(key, uint64(increment)))
	}

}

func (sh *StreamHeap) Top() []Result {
	results := []Result{}
	for {
		node := sh.heap.Pop()
		if node == nil {
			break
		}
		results = append(results, Result{Key: node.Key, Count: node.Counter})
	}
	n := len(results)
	for i := 0; i < n/2; i++ {
		results[i], results[n-i-1] = results[n-i-1], results[i]
	}
	return results
}
