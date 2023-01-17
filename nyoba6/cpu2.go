package nyoba6

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func Cpu2() {
	fmt.Println(runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(16)
	wg.Add(8)

	startTime := time.Now()

	go countaWg() // pass out a badge to each worker
	go countbWg()
	go countcWg()
	go countdWg()
	go counteWg()
	go countfWg()
	go countgWg()
	go counthWg()

	wg.Wait()

	elapsed := time.Since(startTime)
	fmt.Printf("Processes took %s", elapsed)
}

func countaWg() {
	fmt.Println("AAAA is starting  ")
	for I := 1; I < 10_000_000_000; I++ {
	}

	fmt.Println("AAAA is done  ")
	wg.Done() // Turn in my badge - I'm done

}
func countbWg() {
	fmt.Println("BBBB is starting  ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	fmt.Println("BBBB is done")
	wg.Done()

}
func countcWg() {
	fmt.Println("CCCC is starting  ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	fmt.Println("CCCC is done    ")
	wg.Done()

}
func countdWg() {
	fmt.Println("DDDD is starting  ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	fmt.Println("DDDD is done     ")
	wg.Done()

}
func counteWg() {
	fmt.Println("EEEE is starting  ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	fmt.Println("EEEE is done   ")
	wg.Done()

}
func countfWg() {
	fmt.Println("FFFF is starting  ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	fmt.Println("FFFF is done     ")
	wg.Done()

}
func countgWg() {
	fmt.Println("GGGG is starting  ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	fmt.Println("GGGG is done     ")
	wg.Done()

}
func counthWg() {
	fmt.Println("HHHH is starting  ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	fmt.Println("HHHH is done     ")
	wg.Done()

}
