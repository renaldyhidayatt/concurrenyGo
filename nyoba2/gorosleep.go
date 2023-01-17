package nyoba2

import (
	"fmt"
	"time"
)

func GoroSleep() {
	start := time.Now()
	go doSomething()
	go doSomethingElse()

	time.Sleep(time.Second * 5)

	fmt.Println("\n\nI guess I'm done")
	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}

func doSomethingGoSleep() {
	time.Sleep(time.Second * 2)
	fmt.Println("\nI've done something")
}

func doSomethingElseGoSleep() {
	time.Sleep(time.Second * 2)
	fmt.Println("I've done something else")
}
