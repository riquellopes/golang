package main

import (
	"fmt"
	"sync"
)

// https://acervolima.com/mutex-em-golang-com-exemplos/

var GFG = 0

func worker(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()

	GFG = GFG + 1
	m.Unlock()

	wg.Done()
}

func main() {
	var w sync.WaitGroup

	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		w.Add(1)

		go worker(&w, &m)
	}

	w.Wait()

	fmt.Println("Value of x", GFG)
}
