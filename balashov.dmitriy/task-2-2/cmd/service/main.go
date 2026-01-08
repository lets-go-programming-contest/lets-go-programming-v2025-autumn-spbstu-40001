package main

import (
	"fmt"
	"runtime"
	"time"
)

const (
	objectsCount = 100_000
	objectSize   = 1024
	mbDivisor    = 1024 * 1024
)

func allocate() {
	data := make([][]byte, objectsCount)
	for i := range data {
		data[i] = make([]byte, objectSize)
	}
}

func main() {
	var stats runtime.MemStats

	runtime.ReadMemStats(&stats)
	fmt.Printf(
		"Before allocation — Heap: %v MB, GC cycles: %v\n",
		stats.HeapAlloc/mbDivisor,
		stats.NumGC,
	)

	allocate()

	runtime.ReadMemStats(&stats)
	fmt.Printf(
		"After allocation  — Heap: %v MB, GC cycles: %v\n",
		stats.HeapAlloc/mbDivisor,
		stats.NumGC,
	)

	fmt.Println("Calling runtime.GC()...")

	gcStart := time.Now()

	runtime.GC()

	fmt.Printf("GC duration: %v\n", time.Since(gcStart))

	runtime.ReadMemStats(&stats)
	fmt.Printf(
		"After GC          — Heap: %v MB, GC cycles: %v\n",
		stats.HeapAlloc/mbDivisor,
		stats.NumGC,
	)

	time.Sleep(time.Second)
}
