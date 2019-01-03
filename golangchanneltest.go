package main 

import ("fmt"
"os"
)

func main() {
	ch := make(chan int)

	fmt.Println("\nStarting....")
	fmt.Println("======")

	for x,_ := range os.Args {
		fmt.Printf("Parm [%d] : [%s]\n", x, os.Args[x] )
	}

	fmt.Println("======")
	go func(){
		for i := 0;i<10; i++ {
			fmt.Printf("[out: %d] ", i)
			ch <- i
			//time.Sleep(time.Second)
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Printf("[in: %d] ", i)
	}

	fmt.Println("\nComplete.\n")
}
