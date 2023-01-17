package nyoba6

import (
	"fmt"
	"runtime"
	"time"
)

func CpuChannel() {
	fmt.Println(runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(16)
	c := make(chan string)
	startTime := time.Now()
	go countaChan(c)
	go countbChan(c)
	go countcChan(c)
	go countdChan(c)
	go counteChan(c)
	go countfChan(c)
	go countgChan(c)
	go counthChan(c)

	for i := 0; i < 8; i++ {
		fmt.Println(<-c)
	}

	elapsed := time.Since(startTime)
	fmt.Printf("Processes took %s", elapsed)
}

func countaChan(c chan string) {
	fmt.Println("AAAA is starting  ")
	for I := 1; I < 10_000_000_000; I++ {
	}

	c <- "AAAA is done"

}
func countbChan(c chan string) {
	fmt.Println("BBBB is starting  ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	c <- "BBBB is done"

}
func countcChan(c chan string) {
	fmt.Println("CCCC is starting     ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	c <- "CCCC is done"

}
func countdChan(c chan string) {
	fmt.Println("DDDD is starting     ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	c <- "DDDD is done"

}
func counteChan(c chan string) {
	fmt.Println("EEEE is starting     ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	c <- "EEEE is done"

}
func countfChan(c chan string) {
	fmt.Println("FFFF is starting     ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	c <- "FFFF is done"

}
func countgChan(c chan string) {
	fmt.Println("GGGG is starting     ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	c <- "GGGG is done"

}
func counthChan(c chan string) {
	fmt.Println("HHHH is starting     ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	c <- "HHHH is done"

}
