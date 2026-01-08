package main

import (
	"fmt"
	"runtime"
)

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("HeapAlloc before: %v MB\n", m.HeapAlloc/1024/1024)

	data := make([][]byte, 100000)
	for i := range data {
		data[i] = make([]byte, 1024)
	}

	runtime.ReadMemStats(&m)
	fmt.Printf("HeapAlloc after allocation: %v MB\n", m.HeapAlloc/1024/1024)
}
