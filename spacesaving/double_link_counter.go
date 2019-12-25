package spacesaving

type Counter struct {
	Next         *Counter
	Prev         *Counter
	Value        uint64
	ErrorCount   uint64
	Key          string
	ParentBucket *Bucket
}

type DoubleLinkedCounter struct {
	Head *Counter
	Tail *Counter
}

func NewDoubleLinkedCounter() *DoubleLinkedCounter {
	return &DoubleLinkedCounter{
		Head: nil,
		Tail: nil,
	}
}

func (dlc *DoubleLinkedCounter) insertEnd(node *Counter) {
	if dlc.Tail == nil {
		dlc.insertBeginning(node)
	} else {
		dlc.insertAfter(dlc.Tail, node)
	}
}

func (dlc *DoubleLinkedCounter) insertBeginning(node *Counter) {
	if dlc.Head == nil {
		dlc.Head = node
		dlc.Tail = node
		node.Prev = nil
		node.Next = nil
	} else {
		dlc.insertBefore(dlc.Head, node)
	}
}

func (dlc *DoubleLinkedCounter) insertAfter(node *Counter, newNode *Counter) {
	newNode.Prev = node
	newNode.Next = node.Next
	if node.Next == nil {
		dlc.Tail = newNode
	} else {
		node.Next.Prev = newNode
	}
	node.Next = newNode
}

func (dlc *DoubleLinkedCounter) insertBefore(node *Counter, newNode *Counter) {
	newNode.Prev = node.Prev
	newNode.Next = node
	if node.Prev == nil {
		dlc.Head = newNode
	} else {
		node.Prev.Next = newNode
	}
	node.Prev = newNode
}

func (dlc *DoubleLinkedCounter) Remove(node *Counter) {
	if node.Prev == nil {
		dlc.Head = node.Next
	} else {
		node.Prev.Next = node.Next
	}
	if node.Next == nil {
		dlc.Tail = node.Prev
	} else {
		node.Next.Prev = node.Prev
	}
}
func (dlc *DoubleLinkedCounter) Empty() bool {
	return dlc.Head == nil
}
