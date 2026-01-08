package main

import (
	"fmt"
	"runtime"
	"time"
)

const (
	objectsCount = 100_000
	objectSize   = 1024
)

func main() {
	stats := runtime.MemStats{}

	runtime.ReadMemStats(&stats)
	fmt.Printf("Before allocation — Heap: %v MB, GC cycles: %v\n",
		stats.HeapAlloc/1024/1024, stats.NumGC)

	data := make([][]byte, objectsCount)
	for index := range data {
		data[index] = make([]byte, objectSize)
	}

	runtime.ReadMemStats(&stats)
	fmt.Printf("After allocation — Heap: %v MB, GC cycles: %v\n",
		stats.HeapAlloc/1024/1024, stats.NumGC)

	fmt.Println("Calling runtime.GC()...")
	gcStart := time.Now()
	runtime.GC()
	fmt.Printf("GC duration: %v\n", time.Since(gcStart))

	runtime.ReadMemStats(&stats)
	fmt.Printf("After GC — Heap: %v MB, GC cycles: %v\n",
		stats.HeapAlloc/1024/1024, stats.NumGC)

	_ = data

	time.Sleep(time.Second)
}
