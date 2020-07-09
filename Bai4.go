package main

import (
	"bufio"
	"log"
	"os"
	"sync"
)

type line struct {
	value string
	line  int
}

func PrintLine(chanlbuff chan *line, wg *sync.WaitGroup) {
	for {
		select {
		case dat := <-chanlbuff:
			log.Printf("Dong %d giá trị là: %s xong!", dat.line, dat.value)
			wg.Done()
		}
	}
}

func Bai4() {
	chanlbuff := make(chan *line, 10)
	defer close(chanlbuff)

	var wg sync.WaitGroup

	file, err := os.Open("file.txt")
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < 3; i++ {
		go func() {
			PrintLine(chanlbuff, &wg)

		}()
	}
	i := 1
	input := bufio.NewScanner(file)

	for input.Scan() {
		output := line{value: input.Text(), line: i}
		i++
		chanlbuff <- &output
		wg.Add(1)
	}

	wg.Wait()
}
