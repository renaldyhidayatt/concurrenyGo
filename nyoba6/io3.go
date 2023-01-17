package nyoba6

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func Io3() {
	runtime.GOMAXPROCS(8)

	links := []string{
		"http://dev.to",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://medium.com",
		"http://github.com",
	}
	c := make(chan string, len(links))

	start := time.Now()

	for _, link := range links {
		go checkLinkChan(link, c)
	}
	for len(c) < len(links) {

	}
	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}

func checkLinkChan(link string, c chan string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is not responding!")
		c <- link
		return
	}

	fmt.Println(link, "is LIVE!")
	c <- link
}
