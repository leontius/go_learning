package fetch

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"time"
)

func New(urls ...string) {
	start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		// start a goroutine
		go fetch(url, ch)
	}
	for range urls {
		// receive from channel ch
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		// send error to channel ch
		ch <- fmt.Sprint(err)
		return
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	b, err := httputil.DumpResponse(resp, true)
	if err != nil {
		ch <- fmt.Sprintf("dump response error %s: %v err: %v", url, resp, err)
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d byte  %s \n%s", secs, nbytes, url, string(b))
}
