package task2

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func webFetcher(ch chan string, wg *sync.WaitGroup, url string) {
	defer wg.Done()
	time.Sleep(time.Duration(rand.Intn(5000)+1) * time.Millisecond)
	ch <- url
}

func DisplayWebFetcher() {
	urlLists := []string{
		"http://web1.com",
		"http://web2.com",
		"http://web3.com",
		"http://web4.com",
		"http://web5.com",
	}

	ch := make(chan string, len(urlLists))
	wg := sync.WaitGroup{}

	for _, url := range urlLists {
		wg.Add(1)
		go webFetcher(ch, &wg, url)
	}
	wg.Wait()

	go func() {
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}
}
