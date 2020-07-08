package main

import (
	"log"
	"sync"
	"time"
)

func errFunc() {
	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 1; j < 10000; j++ {

				if _, ok := m[j]; !ok {

					delete(m, j)
					continue

				}

				m[j] = j * 10

			}
		}()
	}
	log.Println(m)
	log.Print("done")
}

func main() {

	//// Bai 1: dùng kiến thức về go routine và chan đề func dưới in ra đủ 3 message
	/// nâng cao. In ra các message theo thứ tự. -- In ra message 3 trước message 2. Sử dụng 3 cách để làm( gợi ý: sử dụng mutex, chan, waitGroup)

	//Dung chan
	d := make(chan bool)
	go chanRoutine(d)
	<-d

	//Dung waitGroup

	var wg sync.WaitGroup
	wg.Add(1)
	go WaitGroupRoutine(&wg)
	wg.Wait()

	// Dung Mutex
	var mu = &sync.Mutex{}
	log.Print("hello 1")
	go func() {
		time.Sleep(1 * time.Second)
		mu.Lock()
		log.Print("hello 3")
		mu.Unlock()
	}()
	time.Sleep(2 * time.Second)
	log.Print("hello 2")

	/// Bai 2: tạo 1 biến X map[string]string và 3 goroutine cùng thêm dữ liệu vào X. Mỗi goroutine thêm 1000 key khác nhau. Sao cho quá trình đủ 15 key không mất mát dữ liệu.
	X()

	time.Sleep(5 * time.Second)

	/// Bai 3: chạy đoạn chương trình dưới đây. Nếu có lỗi hãy thêm logic để nó chạy đúng.
	errFunc()

	/// Bai 4: bài tập worker pool: tạo bằng tay file dưới. file.txt sau đó đọc từng dòng file này nạp dữ liệu vào 1 buffer channel có size 10, Điều kiện đọc file từng dòng. Chỉ được sử dụng 3 go routine. Kết quả xử lý xong ỉn ra màn hình + từ xong

	Bai4()
}
