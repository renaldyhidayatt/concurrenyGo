package nyoba4

import (
	"fmt"
	"time"
)

func WaitGroup2() {
	start := time.Now()
	doSomething2()
	doSomethingElse2()
	fmt.Println("I guess I'm done")
	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}

func doSomething2() {
	time.Sleep(time.Millisecond * 1500)
	fmt.Println("\nI've done something")
}

func doSomethingElse2() {
	time.Sleep(time.Millisecond * 1500)
	fmt.Println("I've done something else")
}
