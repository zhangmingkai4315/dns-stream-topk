package spacesaving

type Node struct {
	next *Node
	prev *Node
	item *Bucket
}

type DoubleLinkedList struct {
	head *Node
	tail *Node
}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{
		head: nil,
		tail: nil,
	}
}

func (dll *DoubleLinkedList) Head() *Node {
	return dll.head
}

func (dll *DoubleLinkedList) Tail() *Node {
	return dll.tail
}

func CreateDoubleLinkedListFromSlice(dataSet []*Bucket) *DoubleLinkedList {
	dll := NewDoubleLinkedList()
	for _, data := range dataSet {
		node := Node{
			item: data,
		}
		dll.insertEnd(&node)
	}
	return dll
}

func (dll *DoubleLinkedList) insertEnd(node *Node) {
	if dll.tail == nil {
		dll.insertBeginning(node)
	} else {
		dll.insertAfter(dll.tail, node)
	}
}

func (dll *DoubleLinkedList) insertBeginning(node *Node) {
	if dll.head == nil {
		dll.head = node
		dll.tail = node
	} else {
		dll.insertBefore(dll.head, node)
	}
}

func (dll *DoubleLinkedList) insertAfter(position *Node, node *Node) {
	node.prev = position
	node.next = position.next
	if position.next == nil {
		dll.tail = node
	} else {
		position.next.prev = node
	}
	position.next = node
}

func (dll *DoubleLinkedList) insertBefore(position *Node, node *Node) {
	node.prev = position.prev
	node.next = position

	if position.prev == nil {
		dll.head = node
	} else {
		position.prev.next = node
	}
	position.prev = node
}

func (dll *DoubleLinkedList) Remove(node *Node) {
	if node.prev == nil {
		dll.head = node.next
	} else {
		node.prev.next = node.next
	}
	if node.next == nil {
		dll.tail = node.prev
	} else {
		node.next.prev = node.prev
	}
}
