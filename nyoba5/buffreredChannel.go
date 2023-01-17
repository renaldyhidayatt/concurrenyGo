package nyoba5

import "fmt"

func BuffredChannel() {
	c := make(chan string, 3)
	c <- "Hello "
	c <- "Earth "
	c <- "from Mars"

	msg := <-c
	fmt.Print(msg)

	msg = <-c
	fmt.Print(msg)

	msg = <-c
	fmt.Print(msg)
}
