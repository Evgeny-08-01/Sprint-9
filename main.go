package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		fmt.Printf("Size %d<=0, Exceeded array size\n", size)
		return []int{}
	}
	s := make([]int, size)
	for i := 0; i < size; i++ {
		s[i] = rand.Int()
	}
	return s
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) < CHUNKS {
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
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	var delta = (len(data) / CHUNKS) + 1
	array := make([]int, CHUNKS)
	var wg sync.WaitGroup
	wg.Add(CHUNKS)
	for i := 0; i < CHUNKS; i++ {

		start := delta * i
		stop := start + delta
		if i == CHUNKS-1 {
			stop = len(data)
		}
		dataTemp := data[start:stop]
		go func(i int, dataTemp []int) {
			defer wg.Done()
			max := maximum(dataTemp)
			array[i] = max
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
