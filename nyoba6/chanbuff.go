package nyoba6

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func ChanBuff() {
	runtime.GOMAXPROCS(1)
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string, len(links))

	start := time.Now()

	for _, link := range links {
		go checkLinkChannel(link, c)
	}

	for range links {
		fmt.Println("channel message:", <-c)
	}

	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}

func checkLinkChannel(link string, c chan string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
