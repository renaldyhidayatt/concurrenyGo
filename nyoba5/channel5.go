package nyoba5

import (
	"fmt"
	"sync"
	"time"
)

func ChannelSync5() {
	var wg sync.WaitGroup

	for i := 1; i <= 100000; i++ {
		wg.Add(1)
		go func(j int) {

			var result = 0
			goChan := make(chan int)
			mainChan := make(chan string)
			calculateSquare := func() {
				time.Sleep(time.Second * 3)
				result = j * j
				goChan <- result
			}
			reportResult := func() {
				fmt.Println(j, "squared is", <-goChan)

				mainChan <- "You can quit now.  I'm done."
			}

			go calculateSquare()
			go reportResult()
			<-mainChan
			wg.Done()
		}(i)
	}

	wg.Wait()
}
