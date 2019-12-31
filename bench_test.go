package main

import (
	"time"
	"testing"
	"github.com/zhangmingkai4315/stream-topk/spacesaving"
	"github.com/zhangmingkai4315/stream-topk/stream"
)

func makeFlow() (chan string) {
	config := map[string]int{
		"a.com": 20,
		"b.com": 17,
		"c.com": 15,
		"d.com": 13,
		"e.com": 12,
		"f.com": 11,
		"g.com": 10,
		"h.com": 9,
		"i.com": 8,
		"j.com": 7,
		"k.com": 6,
		"l.com": 5,
		"m.com": 3,
	}
	flow, err := stream.NewDNSFlow(config, 0, time.Duration(10)*time.Second)
	if err != nil {
		panic(err)
	}
	stream := flow.Start()
	return stream
}

func BenchmarkSsHeap(b *testing.B) {
	stream := makeFlow()
	manager := spacesaving.NewStreamHeap(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {	
		domain := <-stream
		manager.Offer(domain, 1)
	}
}

func BenchmarkSummary(b *testing.B) {
	stream := makeFlow()
	manager := spacesaving.NewStreamSummary(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {	
		domain := <-stream
		manager.Offer(domain, 1)
	}
}

func BenchmarkFss(b *testing.B) {
	stream := makeFlow()
	manager := spacesaving.NewFilterSpaceSaving(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {	
		domain := <-stream
		manager.Offer(domain, 1)
	}
}