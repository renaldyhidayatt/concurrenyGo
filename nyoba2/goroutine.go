package nyoba2

import (
	"fmt"
	"time"
)

func GoroutineDoang() {
	start := time.Now()
	go doSomething()
	go doSomethingElse()

	fmt.Println("\n\nI guess I'm done")
	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}

func doSomethingGo() {
	time.Sleep(time.Second * 2)
	fmt.Println("\nI've done something")
}

func doSomethingElseGo() {
	time.Sleep(time.Second * 2)
	fmt.Println("I've done something else")
}
