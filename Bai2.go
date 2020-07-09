package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// Des asd asd
type Des struct {
	sync.Mutex
	m    map[string]string
	iden int
}

// Bai2 asd asd
func Bai2() {
	newDes := &Des{m: make(map[string]string), iden: 1}

	for i := 0; i < 3; i++ {
		go func() {
			for i := 0; i < 1000; i++ {

				newDes.Match("Key", "Value")

			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println(newDes.m)
}

// Match method
func (a *Des) Match(key string, value string) {
	a.Lock()
	keyN := key + strconv.FormatInt(int64(a.iden), 10)
	valueN := value + strconv.FormatInt(int64(a.iden), 10)
	a.iden++
	defer a.Unlock()
	a.m[keyN] = valueN
}
