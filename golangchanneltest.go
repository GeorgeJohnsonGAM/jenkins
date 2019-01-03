package main 

import ("fmt"
"os"
)

func main() {
	ch := make(chan int)

	fmt.Println("\nStarting....")
	fmt.Println("======")

	if len(os.Args) == 2 {
		fmt.Printf("ENV VAR [%s]:[%s]\n", os.Args[1], os.Getenv(os.Args[1]))
	}

	// for x,_ := range os.Args {
	// 	fmt.Printf("Parm [%d] : [%s]\n", x, os.Args[x] )
	// }

	fmt.Println("======")
	fmt.Printf("ENV VAR [%s]:[%s]\n", "MY_SECRET", os.Getenv("MY_SECRET"))
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
