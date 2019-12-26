package main

import (
	"fmt"
	"time"

	"github.com/zhangmingkai4315/stream-topk/spacesaving"
	"github.com/zhangmingkai4315/stream-topk/stream"
)

func main() {
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

	for {
		streamManagers := []spacesaving.StreamManager{
			spacesaving.NewStreamSummary(10),
			spacesaving.NewStreamHeap(10),
		}
		for _, manager := range streamManagers {
			flow, err := stream.NewDNSFlow(config, 0, time.Duration(10)*time.Second)
			if err != nil {
				panic(err)
			}
			fmt.Printf("start the domain stream generator\n")
			stream := flow.Start()
			count := 0
			for domain := range stream {
				count++
				manager.Offer(domain, 1)
			}
			fmt.Printf("stream generate %d domains per seconds\n", count/10)
			fmt.Printf("----Result of TopK for %s----\n", manager.Name())
			for index, result := range manager.Top() {
				fmt.Printf("[%d]%s, %d \n", index+1, result.Key, result.Count)
			}
		}
	}

}

// ➜  stream-topk (master) ✗ go run main.go
// start the domain stream generator
// stream generate 458780 domains per seconds
// ----Result of TopK----
// [1]baidu.com, 2293900
// [2]google.com, 458780
// [3]cnnica.cn, 344085
// [4]cnnic.cn, 344085
// [5]cnnicc.cn, 344085
// [6]alibaba.com, 114695
// [7]aws.com, 114695
// [8]sinabv.com.cn, 114695
// [9]sina.com.cn, 114695
// [10]sinac.com.cn, 114695
// start the domain stream generator
// stream generate 1538216 domains per seconds
// ----Result of TopK----
// [1]baidu.com, 7691081
// [2]sinabv.com.cn, 384554
// [3]aws.com, 384554
// [4]sina.com.cn, 384554
// [5]awsa.com, 384554
// [6]sinac.com.cn, 384554
// [7]awse.com, 384554
// [8]cnnica.cn, 1153662
// [9]cnnic.cn, 1153662
// [10]cnnicc.cn, 1153662
// [11]google.com, 1538216
// [12]alibaba.com, 384554
