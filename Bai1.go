package main

import (
	"log"
	"sync"
	"time"
)

func WaitGroupRoutine(wg *sync.WaitGroup) {
	log.Print("hello 1")
	wg.Add(1)
	go func() {
		time.Sleep(1 * time.Second)
		log.Print("hello 3")
		wg.Done()
	}()
	wg.Wait()
	log.Print("hello 2")

}

func chanRoutine() {
	done := make(chan bool)
	log.Print("hello 1")

	go func() {
		time.Sleep(time.Second)
		log.Print("hello 3")
		done <- true
	}()
	<-done
	log.Print("hello 2")

}
