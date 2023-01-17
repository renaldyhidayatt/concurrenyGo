package nyoba7

import (
	"fmt"
	"sync"
)

var (
	wg5              sync.WaitGroup
	mutex5                 = sync.Mutex{}
	widgetInventory5 int32 = 1000
	newPurchase            = sync.NewCond(&mutex5)
)

func RaceCondition5() {
	fmt.Println("Starting inventory count = ", widgetInventory5)
	wg5.Add(2)
	go makeSales5()
	go newPurchases5()
	wg5.Wait()
	fmt.Println("Ending inventory count = ", widgetInventory5)
}

func makeSales5() { // 1000000 widgets sold
	for i := 0; i < 3000; i++ {
		mutex5.Lock()
		if widgetInventory5-100 < 0 {
			newPurchase.Wait()
		}
		widgetInventory5 -= 100
		fmt.Println(widgetInventory5)
		mutex5.Unlock()
	}

	wg5.Done()
}

func newPurchases5() { // 1000000 widgets purchased
	for i := 0; i < 3000; i++ {
		mutex5.Lock()
		widgetInventory5 += 100
		fmt.Println(widgetInventory5)
		newPurchase.Signal()
		mutex5.Unlock()
	}
	wg5.Done()
}
