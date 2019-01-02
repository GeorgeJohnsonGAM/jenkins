package main 

import ("fmt"
"time"
)

func main() {
	ch := make(chan int)

go func(){
		for i := 0;i<10; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			//time.Sleep(time.Second)
		}
		close(ch)
	}()

	time.Sleep(time.Second)

	for i := range ch {
		fmt.Printf("received %d\n", i)
	}
}
