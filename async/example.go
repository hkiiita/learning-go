package async

import "time"

func secondaryWork() chan int {
	notifier := make(chan int)

	go func() {
		time.Sleep(10 * time.Second)
		notifier <- 1
	}()

	return notifier
}

// Execute Executes a trivial task
func Execute() {
	println("Beginning to work....")
	notifier := secondaryWork()
	//work
	time.Sleep(5 * time.Second)
	println("Main work done....")
	// waiting for the signalling for the completion of the secondary work.
	<-notifier
	println("Secondry work completed too.....")
	time.Sleep(1 * time.Second)
	println("Exit")

}
