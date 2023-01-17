package nyoba7

import (
	"fmt"
	"sync"
)

var (
	wg6              sync.WaitGroup
	widgetInventory6 int32 = 1000
	c                      = make(chan int32, 6000)
)

func RaceCondition6() {
	fmt.Println("Starting inventory count = ", widgetInventory6)
	wg6.Add(2)
	go makeSales6(c)
	go newPurchases6(c)
	wg6.Wait()
	for len(c) > 0 {
		widgetInventory6 += <-c

	}
	fmt.Println("Ending inventory count = ", widgetInventory6)
}

func makeSales6(c chan int32) { // 300000 widgets sold
	for i := 0; i < 3000; i++ {
		c <- -100
	}

	wg6.Done()
}

func newPurchases6(c chan int32) {
	for i := 0; i < 3000; i++ {
		c <- 100 // Put +100 purchase transaction into the channel
	}
	wg6.Done()
}
