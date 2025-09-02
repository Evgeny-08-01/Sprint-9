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

var mu sync.Mutex

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь

	if size <= 0 {
		fmt.Printf("значение %d<=0, выходит за пределы допустимого размера массива\n", size)
		return []int{}
	}
	s := make([]int, size)
	rand.NewSource(time.Now().UnixNano())
		for i := 0; i < size; i++ {
		s[i] = rand.Intn(size + 1)
	}
	return s
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		panic("Ошибка: пустой срез")
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
	// ваш код здесь
	array := make([]int, CHUNKS)
	var delta = (len(data) / CHUNKS) + 1
	if len(data) == 0 {
		panic("Ошибка: пустой срез")
	}
	if CHUNKS <= 0 {
		panic("Ошибка: количество потоков<=0")
	}

	var wg sync.WaitGroup

	wg.Add(CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		go func(i int) {
			defer wg.Done()
			start := delta * i
			stop := delta * (i + 1)
			if i == CHUNKS-1 {
				stop = len(data)
			}
			max := data[start]
			for j := start; j < stop; j++ {
				if max < data[j] {
					max = data[j]
				}
			}
			mu.Lock()
			array[i] = max
			mu.Unlock()
		}(i)

	}
	wg.Wait()
	return maximum(array)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	arrSIZE := generateRandomElements(SIZE)
	fmt.Println("Ищем максимальное значение в один поток")
	start1 := time.Now()
	max1 := maximum(arrSIZE)
	elapsed := time.Since(start1)
	elapsed1 := elapsed.Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max1, elapsed1)

	fmt.Printf("Ищем максимальное значение в количестве потоков равном %d\n", CHUNKS)
	// ваш код здесь
	start2 := time.Now()
	max2 := maxChunks(arrSIZE)
	elapsed = time.Since(start2)
	elapsed2 := elapsed.Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d µs\n", max2, elapsed2)
}
