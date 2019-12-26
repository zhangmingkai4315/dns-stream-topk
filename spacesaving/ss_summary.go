package spacesaving

type StreamSummary struct {
	capacity      int
	topK          int
	streamCounter uint64

	cache   map[string]*Counter
	buckets *DoubleLinkedBucket
}

func NewStreamSummary(topk int) *StreamSummary {
	bucket := NewBucket(0)
	capacity := topk * 2
	for i := 0; i < capacity; i++ {
		counter := Counter{
			ParentBucket: bucket,
		}
		bucket.Children.insertEnd(&counter)
	}
	buckets := NewDoubleLinkedBucket()
	buckets.insertBeginning(bucket)
	ss := StreamSummary{
		capacity: capacity,
		topK:     topk,
		buckets:  buckets,
		cache:    make(map[string]*Counter),
	}
	return &ss
}
func (ss *StreamSummary) Name() string {
	return "StreamSummary"
}
func (ss *StreamSummary) Offer(item string, increment int) {
	ss.streamCounter++
	if itemInCache, ok := ss.cache[item]; ok == true {
		counter := itemInCache
		ss.incrementCounter(counter, increment)
	} else {
		minElement := ss.buckets.Tail.Children.Head
		// originalMin := minElement.Value
		delete(ss.cache, minElement.Key)
		ss.cache[item] = minElement
		minElement.Key = item
		ss.incrementCounter(minElement, increment)
		// if len(ss.cache) > ss.capacity {
		// 	minElement.ErrorCount = originalMin
		// }
	}
}

func (ss *StreamSummary) incrementCounter(counter *Counter, increment int) {
	bucket := counter.ParentBucket
	bucketNext := bucket.Prev
	bucket.Children.Remove(counter)
	// next := bucket.
	counter.Value += uint64(increment)
	if bucketNext != nil && counter.Value == bucketNext.Counter {
		bucketNext.Children.insertEnd(counter)
		counter.ParentBucket = bucketNext
	} else {
		newBucket := NewBucket(counter.Value)
		newBucket.Children.insertEnd(counter)
		ss.buckets.insertBefore(bucket, newBucket)
		counter.ParentBucket = newBucket
	}
	if bucket.Children.Empty() {
		ss.buckets.Remove(bucket)
	}
}

type Result struct {
	Key   string
	Count uint64
}

func (ss *StreamSummary) Top() []Result {
	results := []Result{}
	currentBucket := ss.buckets.Head
	start := currentBucket.Children.Head
	for start != nil {
		result := Result{
			Key:   start.Key,
			Count: start.Value,
		}
		results = append(results, result)
		if len(results) >= ss.topK {
			return results
		}
		if start.Next == nil {
			if currentBucket.Next != nil && currentBucket.Next.Counter > 0 {
				start = currentBucket.Next.Children.Head
				currentBucket = currentBucket.Next
			} else {
				break
			}
		} else {
			start = start.Next
		}
	}
	return results
}
