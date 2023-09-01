package goroutines

import "fmt"

func concurrency() {
	// channels
	// ch := make(chan string, 1) // 1 indicates buffer, if omitted then unbuffered
	ch := make(chan string, 1)
	go func() {
		for _, word := range []string{"hello", "world"} {
			ch <- word
			close(ch)
		}
	}()
	for word := range ch {
		fmt.Println(word)
	}
}
