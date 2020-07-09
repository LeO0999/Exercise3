package main

import (
	"log"
	"sync"
	"time"
)

func errFunc() {
	mu := new(sync.Mutex)
	m := make(map[int]int)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 1; j < 10000; j++ {
				mu.Lock()
				_, ok := m[j]
				mu.Unlock()
				if ok {
					mu.Lock()
					delete(m, j)
					mu.Unlock()

					continue
				}
				mu.Lock()
				m[j] = j * 10
				mu.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)

	log.Println(m)
	log.Print("done")

}
