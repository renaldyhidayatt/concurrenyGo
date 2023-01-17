package nyoba6

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
)

var wgIo2 = sync.WaitGroup{}

func Io2() {
	runtime.GOMAXPROCS(16) // Even ONE processor can take advantage of concurrency with IO bound code. More may not be needed.

	links := []string{
		"http://hashnode.com",
		"http://dev.to",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://medium.com",
		"http://github.com",
		"http://techcrunch.com",
		"http://techrepublic.com",
	}
	wgIo2.Add(len(links))

	start := time.Now()

	for _, link := range links {
		go checkLinkIo2(link)
	}
	wgIo2.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}

func checkLinkIo2(link string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is not responding!")
		wgIo2.Done()
		return
	}

	fmt.Println(link, "is LIVE!")
	wgIo2.Done()
}
