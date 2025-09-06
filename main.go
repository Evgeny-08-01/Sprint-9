package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int64) []int64 {
	if size <= 0 {
		fmt.Printf("Size %d<=0, Exceeded array size\n", size)
		return []int64{}
	}
	s := make([]int64, size)
	for i := int64(0); i < size; i++ {
		s[i] = int64(rand.Int())
	}
	return s
}

// maximum returns the maximum number of elements.
func maximum(data []int64) int64 {
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for i := range data {
		if max < data[i] {
			max = data[i]
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int64) int64 {
	if len(data) == 0 {
		return 0
	}
	var delta = (len(data) / CHUNKS) + 1
	array := make([]int64, CHUNKS)
	var wg sync.WaitGroup
	wg.Add(CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		start := delta * i
		stop := start + delta
		if i == CHUNKS-1 {
			stop = len(data)
		}
		dataTemp := data[start:stop]
		go func(i int, dataTemp []int64) {
			defer wg.Done()
			max := maximum(dataTemp)
			atomic.StoreInt64(&array[i], int64(max))
		}(i, dataTemp)
	}
	wg.Wait()
	return maximum(array)
}

func main() {
	fmt.Printf("Generating %d integers\n", SIZE)
	arrSIZE := generateRandomElements(SIZE)
	fmt.Println("looking for the maximum value in one stream")
	start1 := time.Now()
	max1 := maximum(arrSIZE)
	elapsed1 := time.Since(start1).Microseconds()
	fmt.Printf("Maximum value of the element: %d\nSearch time: %d ms\n", max1, elapsed1)
	fmt.Printf("looking for the maximum value in the number of streams %d\n", CHUNKS)
	start2 := time.Now()
	max2 := maxChunks(arrSIZE)
	elapsed2 := time.Since(start2).Microseconds()
	fmt.Printf("Maximum value of the element: %d\nSearch time: %d µs\n", max2, elapsed2)
}
