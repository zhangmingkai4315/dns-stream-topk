package spacesaving

type Bucket struct {
	value    int
	children *LinkList
}

func NewBucket(value int) *Bucket {
	linklist := LinkList{head: nil}
	return &Bucket{
		value:    value,
		children: &linklist,
	}
}

func (bucket *Bucket) GetChildren() *LinkList {
	return bucket.children
}

func (bucket *Bucket) SetChildren(ll *LinkList) {
	bucket.children = ll
}

func (bucket *Bucket) SetValue(value int) {
	bucket.value = value
}
