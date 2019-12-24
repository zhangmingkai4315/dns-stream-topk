package spacesaving

type LinkListNode struct {
	next    *LinkListNode
	ID      string
	Epsilon int
}

type LinkList struct {
	head *LinkListNode
}
