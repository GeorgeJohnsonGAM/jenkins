package main 

import ("fmt"
"time"
"os"
)

func main() {
	ch := make(chan int)

	fmt.Println("Starting....")
	
	fmt.Printf("Leaked password : [%s]\n", os.Args[1] )

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

	fmt.Println("Complete.")
}
