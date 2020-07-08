package main

import (
	"fmt"
	"strconv"
	"sync"
)

func X() {
	key := "key"
	value := "value"

	iden := 0
	X := make(map[string]string)
	for i := 0; i < 3; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				var mux = &sync.Mutex{}

				mux.Lock()
				keyN := key + strconv.FormatInt(int64(iden), 10)
				valueN := value + strconv.FormatInt(int64(iden), 10)
				iden++
				defer mux.Unlock()
				X[keyN] = valueN
				fmt.Println(X)
			}
		}()
	}
}
