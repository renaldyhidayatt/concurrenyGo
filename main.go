package main

import (
	"fmt"
	// "simple-concurrency/nyoba1"
	"simple-concurrency/nyoba2"
)

func main() {
	nyoba2.NotUsing()
	fmt.Println("--------------------")
	nyoba2.GoroutineDoang()
	fmt.Println("--------------------")
	nyoba2.GoroSleep()
	// nyoba1.ConcurrencyNyoba()
	// nyoba1.Simple()
}
