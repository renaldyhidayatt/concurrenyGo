package nyoba5

import (
	"fmt"
	"time"
)

func Channel() {
	ch := make(chan string)

	go sendMe(ch)

	fmt.Println(<-ch)
}

func sendMe(ch chan<- string) {

	time.Sleep(time.Second * 2)
	ch <- "SendMe is done"

}
