package main

import (
	"fmt"
)

func chunkSlice[T any](slice []T, chunkSize int) [][]T {
	var chunks [][]T
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

func main() {
	n := 13
	var s []int

	for i := 0; i < n; i++ {
		s = append(s, i)
	}

	chunks := chunkSlice(s, 3)
	for _, chunk := range chunks {
		fmt.Println(chunk)
	}
}
