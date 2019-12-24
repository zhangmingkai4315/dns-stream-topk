package spacesaving

type StreamSummary struct {
	capacity int
	total    uint64

	cache   map[string]*Counter
	buckets *DoubleLinkedList
}
