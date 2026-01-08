package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var m runtime.MemStats

	runtime.ReadMemStats(&m)
	fmt.Printf("Before alloc — Heap: %v MB, GC cycles: %v\n",
		m.HeapAlloc/1024/1024, m.NumGC)

	data := make([][]byte, 100000)
	for i := range data {
		data[i] = make([]byte, 1024)
	}

	runtime.ReadMemStats(&m)
	fmt.Printf("After alloc — Heap: %v MB, GC cycles: %v\n",
		m.HeapAlloc/1024/1024, m.NumGC)

	fmt.Println("Calling runtime.GC()...")
	start := time.Now()
	runtime.GC()
	fmt.Printf("GC took: %v\n", time.Since(start))

	runtime.ReadMemStats(&m)
	fmt.Printf("After GC — Heap: %v MB, GC cycles: %v\n",
		m.HeapAlloc/1024/1024, m.NumGC)

	_ = data
	time.Sleep(1 * time.Second)
}
