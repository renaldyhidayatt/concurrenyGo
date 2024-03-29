package nyoba6

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func Io1() {
	runtime.GOMAXPROCS(16)
	fmt.Println(runtime.NumCPU())

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

	start := time.Now()

	for _, link := range links {
		checkLink(link)
	}

	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}

func checkLink(link string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is not responding!")

		return
	}

	fmt.Println(link, "is LIVE!")

}
