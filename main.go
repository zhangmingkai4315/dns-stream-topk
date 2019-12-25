package main

import (
	"fmt"

	"github.com/zhangmingkai4315/stream-topk/spacesaving"
)

func main() {
	// config := map[string]int{
	// 	"baidu.com":     20,
	// 	"google.com":    4,
	// 	"cnnic.cn":      3,
	// 	"sina.com.cn":   1,
	// 	"aws.com":       1,
	// 	"cnnica.cn":     3,
	// 	"sinabv.com.cn": 1,
	// 	"awse.com":      1,
	// 	"cnnicc.cn":     3,
	// 	"sinac.com.cn":  1,
	// 	"awsa.com":      1,
	// 	"alibaba.com":   1,
	// }
	// flow, err := stream.NewDNSFlow(config, 0, time.Duration(2)*time.Second)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("start the domain stream generator\n")
	// stream := flow.Start()
	// count := 0
	// summary := spacesaving.NewStreamSummary(10)
	// for domain := range stream {
	// 	count++
	// 	summary.Offer(domain, 1)
	// }
	// fmt.Printf("stream generate %d domains per seconds\n", count/10)
	// fmt.Println("----Result of TopK----")
	// for index, result := range summary.Top() {
	// 	fmt.Printf("[%d]%s, %d \n", index+1, result.Key, result.Count)
	// }

	heap := spacesaving.NewHeap(10)
	for i := 0; i < 20; i++ {
		heap.Insert(spacesaving.NewHeapNode(fmt.Sprintf("abc-%d", i), 40-i))
		fmt.Printf("%v\n", heap.Data)
	}

}
