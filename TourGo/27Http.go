package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	start := time.Now()
	//    ch := make(chan string)
	url := ""

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go fetch(url)
	}

	wg.Wait()

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string) {
	defer wg.Done()

	client := http.DefaultClient

	//提交请求
	req, err := http.NewRequest("GET", url, nil)

	//增加header选项
	req.Header.Add("x-token", "123")

	if err != nil {
		fmt.Printf(fmt.Sprint(err))
		return
	}
	q := req.URL.Query()
	q.Add("version", "0.1")
	q.Add("query", "123")
	req.URL.RawQuery = q.Encode()

	start := time.Now()
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(fmt.Sprint(err))
		return
	}
	defer resp.Body.Close()
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)

	if err != nil {
		fmt.Println(fmt.Sprintf("while reading %s: %v", url, err))
		return
	}
	secs := time.Since(start).Seconds()
	fmt.Println(fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url))
}
