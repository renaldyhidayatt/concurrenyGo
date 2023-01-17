package nyoba7

import (
	"fmt"
	"sync"
)

var (
	wg2              sync.WaitGroup
	mutex2                 = sync.Mutex{}
	widgetInventory2 int32 = 1000
)

func RaceCondition2() {
	fmt.Println("Starting inventory count = ", widgetInventory)
	wg2.Add(2)
	go makeSales2()
	go newPurchases2()
	wg2.Wait()
	fmt.Println("Ending inventory count = ", widgetInventory)
}

func makeSales2() {
	for i := 0; i < 300000; i++ {
		mutex2.Lock()
		widgetInventory2 -= 100
		mutex2.Unlock()
	}

	wg2.Done()
}

func newPurchases2() {
	for i := 0; i < 300000; i++ {
		mutex2.Lock()
		widgetInventory += 100
		mutex2.Unlock()
	}
	wg2.Done()
}
