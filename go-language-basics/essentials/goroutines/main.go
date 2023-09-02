package goroutines 

import(
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		// add one record to WaitGroup counter
		wg.Add(1)
		// anonymous concurrent function
		go func(i int) { 
			// defer == code that executes at the end of block process
			// wg.Done substract 1 from WaitGroup
			defer wg.Done()
			fmt.Println(i)
		}(i) // passing i to func
	}
 e
	fmt.Println("hello")
	// Wait for all goroutines to finish before proceeding
	wg.Wait()
	concurrency()
}
