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
		return []int{}
	}
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(1_000_000_000) + 1
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for _, value := range data[1:] {
		if value > max {
			max = value
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	count := CHUNKS
	if len(data) < count {
		count = len(data)
	}
	maxima := make([]int, count)
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		start, end := i*len(data)/count, (i+1)*len(data)/count
		wg.Add(1)
		go func(index int, chunk []int) {
			defer wg.Done()
			maxima[index] = maximum(chunk)
		}(i, data[start:end])
	}
	wg.Wait()
	return maximum(maxima)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d µs\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d µs\n", max, elapsed)
}
