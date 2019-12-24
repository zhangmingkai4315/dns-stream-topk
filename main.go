package main

import (
	"github.com/zhangmingkai4315/stream-topk/stream"
)

import (
	"fmt"
	"time"
)

func main() {
	config := map[string]int{
		"baidu.com":  1,
		"google.com": 5,
		"amazon.com": 1,
		"bing.com":   1,
	}
	flow, err := stream.NewDNSFlow(config, 0, time.Duration(10)*time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Printf("start the domain stream generator\n")
	stream := flow.Start()
	count := 0
	for range stream {
		count++
	}
	fmt.Printf("stream generate %d domains per seconds\n", count/10)
}
