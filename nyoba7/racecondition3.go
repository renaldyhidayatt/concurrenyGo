package nyoba7

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	wg3 sync.WaitGroup

	widgetInventory3 int32 = 1000 //Package-level variable will avoid all the pointers
)

func RaceCondition3() {
	fmt.Println("Starting inventory count = ", widgetInventory)
	wg3.Add(2)
	go makeSales3()
	go newPurchases3()
	wg3.Wait()
	fmt.Println("Ending inventory count = ", widgetInventory)
}

func makeSales3() {
	for i := 0; i < 300000; i++ {

		atomic.AddInt32(&widgetInventory3, -100)

	}

	wg3.Done()
}

func newPurchases3() {
	for i := 0; i < 300000; i++ {

		atomic.AddInt32(&widgetInventory3, 100)

	}
	wg3.Done()
}
