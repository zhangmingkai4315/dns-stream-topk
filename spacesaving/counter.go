package spacesaving

type Counter struct {
	Value      uint64
	ErrorCount uint64
	Key        string
	Bucket     *DoubleLinkedList
}

func NewCounter(bucket *DoubleLinkedList) *Counter {
	return &Counter{
		Bucket: bucket,
	}
}
