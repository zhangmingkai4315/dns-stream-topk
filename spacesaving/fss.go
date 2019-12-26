package spacesaving

import (
	"container/heap"
	"sort"

	"github.com/dgryski/go-sip13"
)

// Element is a TopK item
type Element struct {
	Key   string
	Count int
	Error int
}

type elementsByCountDescending []Element

func (elts elementsByCountDescending) Len() int { return len(elts) }
func (elts elementsByCountDescending) Less(i, j int) bool {
	return (elts[i].Count > elts[j].Count) || (elts[i].Count == elts[j].Count && elts[i].Key < elts[j].Key)
}
func (elts elementsByCountDescending) Swap(i, j int) { elts[i], elts[j] = elts[j], elts[i] }

type keys struct {
	m    map[string]int
	elts []Element
}

// Implement the container/heap interface

func (tk *keys) Len() int { return len(tk.elts) }
func (tk *keys) Less(i, j int) bool {
	return (tk.elts[i].Count < tk.elts[j].Count) || (tk.elts[i].Count == tk.elts[j].Count && tk.elts[i].Error > tk.elts[j].Error)
}
func (tk *keys) Swap(i, j int) {

	tk.elts[i], tk.elts[j] = tk.elts[j], tk.elts[i]

	tk.m[tk.elts[i].Key] = i
	tk.m[tk.elts[j].Key] = j
}

func (tk *keys) Push(x interface{}) {
	e := x.(Element)
	tk.m[e.Key] = len(tk.elts)
	tk.elts = append(tk.elts, e)
}

func (tk *keys) Pop() interface{} {
	var e Element
	e, tk.elts = tk.elts[len(tk.elts)-1], tk.elts[:len(tk.elts)-1]

	delete(tk.m, e.Key)

	return e
}

// FilterSpaceSaving calculates the TopK elements for a stream
type FilterSpaceSaving struct {
	n      int
	k      keys
	alphas []int
}

// NewFilterSpaceSaving returns a FilterSpaceSaving structrue
func NewFilterSpaceSaving(n int) *FilterSpaceSaving {
	return &FilterSpaceSaving{
		n:      n,
		k:      keys{m: make(map[string]int), elts: make([]Element, 0, n)},
		alphas: make([]int, n*6),
	}
}

// Name return the manager name
func (fss *FilterSpaceSaving) Name() string {
	return "FilterSpaceSaving"
}

func reduce(x uint64, n int) uint32 {
	return uint32(uint64(uint32(x)) * uint64(n) >> 32)
}

// Offer adds an element to the stream to be tracked
func (fss *FilterSpaceSaving) Offer(x string, count int) {
	xhash := reduce(sip13.Sum64Str(0, 0, x), len(fss.alphas))
	// are we tracking this element?
	if idx, ok := fss.k.m[x]; ok {
		fss.k.elts[idx].Count += count
		// e := fss.k.elts[idx]
		heap.Fix(&fss.k, idx)
		return
	}

	// can we track more elements?
	if len(fss.k.elts) < fss.n {
		// there is free space
		e := Element{Key: x, Count: count}
		heap.Push(&fss.k, e)
		return
	}

	if fss.alphas[xhash]+count < fss.k.elts[0].Count {
		//  := Element{
		// 	Key:   x,
		// 	Error: fss.alphas[xhash],
		// 	Count: fss.alphas[xhash] + count,
		// }
		fss.alphas[xhash] += count
		return
	}

	// replace the current minimum element
	minKey := fss.k.elts[0].Key

	mkhash := reduce(sip13.Sum64Str(0, 0, minKey), len(fss.alphas))
	fss.alphas[mkhash] = fss.k.elts[0].Count

	e := Element{
		Key:   x,
		Error: fss.alphas[xhash],
		Count: fss.alphas[xhash] + count,
	}
	fss.k.elts[0] = e
	delete(fss.k.m, minKey)
	fss.k.m[x] = 0
	heap.Fix(&fss.k, 0)
	return
}

// Keys returns the current estimates for the most frequent elements
func (fss *FilterSpaceSaving) Keys() []Element {
	elts := append([]Element(nil), fss.k.elts...)
	sort.Sort(elementsByCountDescending(elts))
	return elts
}

// Estimate returns an estimate for the item x
func (fss *FilterSpaceSaving) Estimate(x string) Element {
	xhash := reduce(sip13.Sum64Str(0, 0, x), len(fss.alphas))

	// are we tracking this element?
	if idx, ok := fss.k.m[x]; ok {
		e := fss.k.elts[idx]
		return e
	}
	count := fss.alphas[xhash]
	e := Element{
		Key:   x,
		Error: count,
		Count: count,
	}
	return e
}

func (fss *FilterSpaceSaving) Top() []Result {
	results := []Result{}
	elts := append([]Element(nil), fss.k.elts...)
	sort.Sort(elementsByCountDescending(elts))
	for _, item := range elts {
		results = append(results, Result{
			Key:   item.Key,
			Count: uint64(item.Count),
		})
	}
	return results
}
