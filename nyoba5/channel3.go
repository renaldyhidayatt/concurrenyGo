package nyoba5

import (
	"fmt"
	"sync"
	"time"
)

func Channel2() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		var value = 97
		var result = 0

		goChan := make(chan int)
		mainChan := make(chan string)
		calculateSquare := func() {
			time.Sleep(time.Second * 3)
			result = value * value
			goChan <- result
		}

		reportResult := func() {
			fmt.Println(value, "squared is", <-goChan)
			// blocks until it can read something from goChan - printed
			mainChan <- "You can quit now.  I'm done."
		}
		go calculateSquare()
		go reportResult()
		<-mainChan
		wg.Done()
	}()
	wg.Done()
}
