package nyoba7

import (
	"fmt"
	"sync"
)

var (
	wg4              sync.WaitGroup
	mutex4                 = sync.Mutex{}
	widgetInventory4 int32 = 1000
)

func RaceCondition4() {
	fmt.Println("Starting inventory count = ", widgetInventory4)
	wg4.Add(2)
	go makeSales4()
	go newPurchases4()
	wg4.Wait()
	fmt.Println("Ending inventory count = ", widgetInventory4)
}

func makeSales4() {
	for i := 0; i < 3000; i++ {
		mutex4.Lock()
		widgetInventory4 -= 100
		fmt.Println(widgetInventory4)
		mutex4.Unlock()
	}

	wg4.Done()
}

func newPurchases4() {
	for i := 0; i < 3000; i++ {
		mutex4.Lock()
		widgetInventory4 += 100
		fmt.Println(widgetInventory4)
		mutex4.Unlock()
	}
	wg4.Done()
}
