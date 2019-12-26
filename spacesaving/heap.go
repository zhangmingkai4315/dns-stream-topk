package spacesaving

import "fmt"

// HeapNode define the heap node
type HeapNode struct {
	Key     string
	Error   uint64
	Counter uint64
}

// Value return the node value
func (node *HeapNode) String() string {
	return fmt.Sprintf("%s:%d(%d)", node.Key, node.Counter, node.Error)
}

// NewHeapNode create a new heap node
func NewHeapNode(value string, counter uint64) *HeapNode {
	return &HeapNode{
		Key:     value,
		Counter: counter,
		Error:   0,
	}
}

//Heap define a MinHeap struct with
//size and structure
type Heap struct {
	Data []*HeapNode
	N    int
	Max  int
}

// NewHeap create Heap object
func NewHeap(maxN int) *Heap {
	return &Heap{
		Data: []*HeapNode{},
		N:    0,
		Max:  maxN,
	}
}

// IsEmpty returns true if this priority heap is empty.
func (heap *Heap) IsEmpty() bool {
	return heap.N == 0
}

// IsFull returns true if this priority heap is full.
func (heap *Heap) IsFull() bool {
	return heap.N == heap.Max
}

// Min return the min value of Data
func (heap *Heap) Min() (*HeapNode, error) {
	if heap.IsEmpty() {
		return nil, nil
	}
	return heap.Data[0], nil
}

func (heap *Heap) BuildFromSlice(data []*HeapNode) {
	n := len(data)
	heap.Data = data
	for i := n/2 - 1; i >= 0; i-- {
		heap.down(i, n)
	}
}

func (heap *Heap) down(index int, size int) bool {
	i := index
	for {
		j1 := 2*i + 1
		if j1 >= size || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < size && heap.Data[j2].Counter < heap.Data[j1].Counter {
			j = j2 // = 2*i + 2
		}
		if !(heap.Data[j].Counter < heap.Data[i].Counter) {
			break
		}
		heap.Data[i], heap.Data[j] = heap.Data[j], heap.Data[i]
		i = j
	}
	return i > index
}

func (heap *Heap) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !(heap.Data[j].Counter < heap.Data[i].Counter) {
			break
		}
		heap.Data[i], heap.Data[j] = heap.Data[j], heap.Data[i]
		j = i
	}
}

// Push new element to heap
func (heap *Heap) Push(node *HeapNode) {
	if heap.N < heap.Max {
		heap.N++
	}
	heap.Data = append(heap.Data, node)
	heap.up(heap.N - 1)
}

// Pop the min element from the heap
func (heap *Heap) Pop() *HeapNode {
	if heap.N == 0 {
		return nil
	}
	n := heap.N - 1
	heap.Data[0], heap.Data[n] = heap.Data[n], heap.Data[0]
	heap.down(0, n)
	heap.N--
	return heap.Data[n]
}

// Fix the heap when the element changed
func (heap *Heap) Fix(i int) {
	if !(heap.down(i, heap.N)) {
		heap.up(i)
	}
}
