package contexts

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	fmt.Printf("FETCHING VALUE FROM PASSED CONTEXT : %s\n", ctx.Value("myKey"))

	//anotherContext := context.WithValue(ctx, "myKey", "AnotherValue")
	anotherContext, cancelCtx := context.WithCancel(ctx)

	printChannel := make(chan int)

	go doAnother(anotherContext, printChannel)

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		printChannel <- i
	}

	cancelCtx()

	fmt.Printf("FETCHING AGAIN FROM PASSED CONTEXT : %s\n", ctx.Value("myKey"))
	time.Sleep(2 * time.Second)
}

func doAnother(ctx context.Context, printChannel chan int) {

	for {
		fmt.Println("INSIDE LOOP..")
		select {

		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Println("ERROR :", err)
			}
			fmt.Println("CONTEXT CANCELLED !!! EXITING !")
			return

		case i := <-printChannel:
			fmt.Println("READING NUMBER CHANNEL : ", i)
		}
	}

}

func LearnContexts() {

	ctx := context.Background()
	ctx = context.WithValue(ctx, "myKey", "myValue")
	doSomething(ctx)
}
