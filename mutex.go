package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

func main3() {
	var wg sync.WaitGroup

	input := []int{1, 2, 3, 4, 5}
	result := []int{}

	for _, data := range input {
		wg.Add(1)

		go processData(&wg, &result, data)
	}

	wg.Wait()

	fmt.Println(result)
}

func processData(wg *sync.WaitGroup, result *[]int, data int) {
	defer wg.Done()

	processedData := data * 2

	lock.Lock()
	*result = append(*result, processedData)

	lock.Unlock()
}
