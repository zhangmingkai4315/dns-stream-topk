package spacesaving

type Bucket struct {
	Next     *Bucket
	Prev     *Bucket
	Counter  uint64
	Children *DoubleLinkedCounter
}

func NewBucket(counter uint64) *Bucket {
	return &Bucket{
		Counter:  counter,
		Children: NewDoubleLinkedCounter(),
	}
}

type DoubleLinkedBucket struct {
	Head *Bucket
	Tail *Bucket
}

func NewDoubleLinkedBucket() *DoubleLinkedBucket {
	return &DoubleLinkedBucket{
		Head: nil,
		Tail: nil,
	}
}

func (dlb *DoubleLinkedBucket) insertEnd(node *Bucket) {
	if dlb.Tail == nil {
		dlb.insertBeginning(node)
	} else {
		dlb.insertAfter(dlb.Tail, node)
	}
}

func (dlb *DoubleLinkedBucket) insertBeginning(node *Bucket) {
	if dlb.Head == nil {
		dlb.Head = node
		dlb.Tail = node
	} else {
		dlb.insertBefore(dlb.Head, node)
	}
}

func (dlb *DoubleLinkedBucket) insertAfter(position *Bucket, node *Bucket) {
	node.Prev = position
	node.Next = position.Next
	if position.Next == nil {
		dlb.Tail = node
	} else {
		position.Next.Prev = node
	}
	position.Next = node
}

func (dlb *DoubleLinkedBucket) insertBefore(position *Bucket, node *Bucket) {
	node.Prev = position.Prev
	node.Next = position

	if position.Prev == nil {
		dlb.Head = node
	} else {
		position.Prev.Next = node
	}
	position.Prev = node
}

func (dlb *DoubleLinkedBucket) Remove(node *Bucket) {
	if node.Prev == nil {
		dlb.Head = node.Next
	} else {
		node.Prev.Next = node.Next
	}
	if node.Next == nil {
		dlb.Tail = node.Prev
	} else {
		node.Next.Prev = node.Prev
	}
}

func (dlb *DoubleLinkedBucket) Empty() bool {
	return dlb.Head == nil
}
