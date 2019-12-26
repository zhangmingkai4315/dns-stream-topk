package utils

import (
	"fmt"
	"runtime"
)

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

type MemoryUsage struct {
	Alloc      uint64
	TotalAlloc uint64
	Sys        uint64
	NumGC      uint32
}

func NewMemoryUsage() *MemoryUsage {
	usage := &MemoryUsage{}
	usage.update()
	return usage
}

func (usage *MemoryUsage) update() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	usage.Alloc = m.Alloc
	usage.TotalAlloc = m.TotalAlloc
	usage.Sys = m.Sys
	usage.NumGC = m.NumGC
}

func (usage *MemoryUsage) DiffPrint() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	if m.Alloc >= usage.Alloc {
		fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc-usage.Alloc))
	} else {
		fmt.Printf("Alloc = -%v MiB", bToMb(usage.Alloc-m.Alloc))
	}
	if m.TotalAlloc >= usage.TotalAlloc {
		fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc-usage.TotalAlloc))
	} else {
		fmt.Printf("\tTotalAlloc = -%v MiB", bToMb(usage.TotalAlloc-m.TotalAlloc))
	}
	if m.Sys >= usage.Sys {
		fmt.Printf("\tSys = %v MiB", bToMb(m.Sys-usage.Sys))
	} else {
		fmt.Printf("\tSys = -%v MiB", bToMb(usage.Sys-m.Sys))
	}
	if m.NumGC >= usage.NumGC {
		fmt.Printf("\tNumGC = %v \n", m.NumGC-usage.NumGC)
	} else {
		fmt.Printf("\tNumGC = -%v \n", usage.NumGC-m.NumGC)
	}
}
func (usage *MemoryUsage) Print() {
	fmt.Printf("Alloc = %v MiB", bToMb(usage.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(usage.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(usage.Sys))
	fmt.Printf("\tNumGC = %v\n", usage.NumGC)
}
