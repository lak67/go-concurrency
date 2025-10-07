package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func mai6n() {
	done := make(chan interface{})
	defer close(done)

	cows := make(chan interface{}, 100)
	pigs := make(chan interface{}, 100)

	go func() {
		for {
			select {
			case <-done:
				return

			case cows <- "moo":
			}
		}
	}()

	go func() {
		for {
			select {
			case <-done:
				return

			case pigs <- "oink":
			}
		}
	}()

	wg.Add(1)
	go consumeCows(done, cows)
	wg.Add(1)
	go consumePigs(done, pigs)

}

func consumePigs(done <-chan interface{}, pigs <-chan interface{}) {
	defer wg.Done()

	for pig := range orDone(done, pigs) {
		// do some complex logic
		fmt.Println(pig)
	}
}

func consumeCows(done <-chan interface{}, cows <-chan interface{}) {
	defer wg.Done()

	for cow := range orDone(done, cows) {
		// do some complex logic
		fmt.Println(cow)
	}
}

func orDone(done, c <-chan interface{}) <-chan interface{} {
	relayStream := make(chan interface{})
	go func() {
		defer close(relayStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if !ok {
					return
				}
				select {
				case relayStream <- v:
				case <-done:
					return
				}
			}
		}
	}()

	return relayStream
}
