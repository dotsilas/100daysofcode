package goroutines

import "fmt"

// channels: receive and gives data to comunicate between goroutines
// ch <- word     ch receive word
// fmt.Println(<- ch)    Println function receives ch data as arg

func concurrency() {
	// channels
	// ch := make(chan string, 1) // 1 indicates buffer, if omitted then unbuffered
	ch := make(chan string, 1)
	go func() {
		// send words to the channel
		for _, word := range []string{"hello", "world"} {
			ch <- word
		}
		// close channel
		close(ch)
	}()

	// receive words from channel
	for word := range ch {
		fmt.Println(word)
	}
}
