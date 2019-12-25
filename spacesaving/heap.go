package spacesaving

import "fmt"

// HeapNode define the heap node
type HeapNode struct {
	Key     string
	Counter int
}

// Value return the node value
func (node *HeapNode) String() string {
	return fmt.Sprintf("%s:%d", node.Key, node.Counter)
}

// NewHeapNode create a new heap node
func NewHeapNode(value string, counter int) *HeapNode {
	return &HeapNode{
		Key:     value,
		Counter: counter,
	}
}

//Heap define a MinHeap struct with
//size and structure
type Heap struct {
	Data []*HeapNode
	N    int
}

// NewHeap create Heap object
func NewHeap(maxN int) *Heap {
	return &Heap{
		Data: make([]*HeapNode, maxN+1),
		N:    0,
	}
}

// IsEmpty returns true if this priority heap is empty.
func (heap *Heap) IsEmpty() bool {
	return heap.N == 0
}

// Size returns the number of keys on this priority heap.
func (heap *Heap) Size() int {
	return heap.N
}

// Min return the min value of Data
func (heap *Heap) Min() (*HeapNode, error) {
	if heap.IsEmpty() {
		return nil, nil
	}
	return heap.Data[1], nil
}

// Insert a new item
func (heap *Heap) Insert(node *HeapNode) {
	if heap.N != len(heap.Data)-1 {
		heap.N++
	}
	heap.Data[heap.N] = node
	heap.swim(heap.N)

}

func (heap *Heap) resize(capcity int) {
	temp := make([]*HeapNode, capcity)
	N := heap.N
	for i := 1; i <= N; i++ {
		temp[i] = heap.Data[i]
	}
	heap.Data = temp
}

// DelMin delete the min priority and return the key value
func (heap *Heap) DelMin() *HeapNode {
	if heap.IsEmpty() {
		return nil
	}
	min := heap.Data[1]
	heap.Exchange(1, heap.N)
	heap.N--
	heap.sink(1)
	if heap.N > 0 && heap.N == (len(heap.Data)-1)/4 {
		heap.resize(len(heap.Data) / 2)
	}
	return min
}

// Exchange i and j items
func (heap *Heap) Exchange(i, j int) {
	heap.Data[i], heap.Data[j] = heap.Data[j], heap.Data[i]
}

// Less return the compare result
func (heap *Heap) Less(i, j int) bool {
	if heap.Data[i].Counter < heap.Data[j].Counter {
		return true
	}
	return false
}

func (heap *Heap) sink(k int) {
	for 2*k <= heap.N {
		j := 2 * k
		if j < heap.N && heap.Less(j, j+1) == false {
			j++
		}
		if heap.Less(k, j) {
			break
		}
		heap.Exchange(k, j)
		k = j
	}
}

func (heap *Heap) swim(k int) {
	for k > 1 && heap.Less(k/2, k) == false {
		heap.Exchange(k, k/2)
		k = k / 2
	}
}
