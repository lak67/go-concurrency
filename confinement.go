package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	input := []int{1, 2, 3, 4, 5}

	result := make([]int, len(input))

	for i, data := range input {
		wg.Add(1)

		go processData1(&wg, &result[i], data)
	}

	wg.Wait()

	fmt.Println(result)
}

func processData1(wg *sync.WaitGroup, resultDest *int, data int) {
	defer wg.Done()

	processedData := data * 2

	*resultDest = processedData
}
