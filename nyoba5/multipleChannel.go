package nyoba5

import (
	"fmt"
	"os"
	"time"
)

func MultipleChannel() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Second)
			c1 <- "Sending every 1 second"

		}
	}()
	go func() {
		for {
			time.Sleep(time.Second * 4)
			c2 <- "Sending every 4 sec"

		}
	}()
	go func() {
		for {
			time.Sleep(time.Second * 10)
			c3 <- "We're done"
		}
	}()

	for {
		select {
		case msg := <-c1:
			fmt.Println(msg)
		case msg := <-c2:
			fmt.Println(msg + " Something cool happened")
		case msg := <-c3:
			fmt.Println(msg)
			os.Exit(0)

		}
	}
}
