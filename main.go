package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	OPTION1  = "Synchronous Execution"
	OPTION2  = "Concurrent Execution"
	OPTION3  = "Concurrent Execution - With all functions as go routines"
	OPTION4  = "Concurrent Execution - With all functions as go routines + sleep function"
	OPTION5  = "Concurrent Execution - Waitgroup"
	OPTION6  = "Concurrent Execution - Channels: Single message"
	OPTION7  = "Concurrent Execution - Channels: All messages (Deadlock)"
	OPTION8  = "Concurrent Execution - Channels: All messages Resolution Method 1"
	OPTION9  = "Concurrent Execution - Channels: All messages Resolution Method 2"
	OPTION10 = "Channel Constraint"
	OPTION11 = "Buffer Channel"
	OPTION12 = "Buffer Channel Deadlock"
	OPTION13 = "Channel Starvation"
	OPTION14 = "Select Statement"
	OPTION15 = "Worker Pool"
)

func main() {

	var choice string
	fmt.Println("1. " + OPTION1)
	fmt.Println("2. " + OPTION2)
	fmt.Println("3. " + OPTION3)
	fmt.Println("4. " + OPTION4)
	fmt.Println("5. " + OPTION5)
	fmt.Println("6. " + OPTION6)
	fmt.Println("7. " + OPTION7)
	fmt.Println("8. " + OPTION8)
	fmt.Println("9. " + OPTION9)
	fmt.Println("10. " + OPTION10)
	fmt.Println("11. " + OPTION11)
	fmt.Println("12. " + OPTION12)
	fmt.Println("13. " + OPTION13)
	fmt.Println("14. " + OPTION14)
	fmt.Println("15. " + OPTION15)
	fmt.Println("Enter your choice : ")
	fmt.Scanln(&choice)
	switch choice {
	case "1":
		fmt.Println("Starting: " + OPTION1)
		/*
			This is an example for synchronous execution
			The sheep will never terminate and the fish will never be printed
		*/
		count("sheep")
		count("fish")
	case "2":
		fmt.Println("Starting: " + OPTION2)
		/*
			Concurrent Execution
			Once we attach the keyword `go` as a prefix to count function, the function will run as a seperate go routine
			Esentially there will be two go routines - main() and the count("sheep")
			This will allow the functions to execute concurrently
		*/
		go count("sheep")
		count("fish")
	case "3":
		fmt.Println("Starting: " + OPTION3)
		/*
			In this execution, we will make all the functions as go routines by attaching `go` to prefix of function call
			This will result in quick exit of the main function
			In go when the main routine finishes, the program exits
			As we are calling 2 seperate go routines (count sheep and count fish), main will fire these 2 go routines and closes it's execution work, thereby closing the program
		*/
		go count("sheep")
		go count("fish")
	case "4":
		fmt.Println("Starting: " + OPTION4)
		/*
			In this execution, we add a time.sleep() for `n` seconds
			This will allow the main to wait for 2 seconds after calling go routines, there by allowing the go routines to execute for `n` seconds
		*/
		go count("sheep")
		go count("fish")
		time.Sleep(time.Second * 2)
	case "5":
		fmt.Println("Starting: " + OPTION5)
		/*
			Use `WaitGroup` to wait finishing the execution till all the subroutines gets finished
			Add to waitgroup to notify that there is a go routine that needs to be monitored
			Call anonymous function and wrapping count function inside it as it is not the responsibility of count to handle WaitGroup
			Call wg.Add(num) to add go routine that needs to be waited
			Call wg.Done() when the execution of go routine is completed
			Call wg.Wait() to have a blocking call till wg.Done is called
		*/
		var wg sync.WaitGroup

		wg.Add(1)

		go func() {
			countUptoN("sheep")
			wg.Done()
		}()

		wg.Wait()
	case "6":
		fmt.Println("Starting: " + OPTION6)
		/*
			Channels can be used to pass the information from the called routine to caller routine
			`<- c` is a blocking call, it will wait till something is receivable
		*/
		c := make(chan string)
		go countWithChannel("sheep", c)
		msg := <-c
		fmt.Println(msg)
	case "7":
		fmt.Println("Starting: " + OPTION7)
		/*
			Channels can be used to pass the information from the called routine to caller routine
			`<- c` is a blocking call, it will wait till something is receivable
			for {} is an infinite loop, it will wait even after the go routine countWithChannel ends
			Go detects this problem in runtime and throws the deadlock
		*/
		c := make(chan string)
		go countWithChannel("sheep", c)
		for {
			msg := <-c
			fmt.Println(msg)
		}
	case "8":
		fmt.Println("Starting: " + OPTION8)
		/*
			Channels can be used to pass the information from the called routine to caller routine
			`<- c` is a blocking call, it will wait till something is receivable
			for {} is an infinite loop, it will wait even after the go routine countWithChannel ends
			Go detects this problem in runtime and throws the deadlock
			To solve this we close the channel - THIS SHOULD BE DONE BY SENDER NOT RECEIVER
		*/
		c := make(chan string)
		go countWithChannelClose("sheep", c)
		for {
			msg, open := <-c

			if !open {
				break
			}

			fmt.Println(msg)
		}
	case "9":
		fmt.Println("Starting: " + OPTION9)
		/*
			Channels can be used to pass the information from the called routine to caller routine
			`<- c` is a blocking call, it will wait till something is receivable
			for {} is an infinite loop, it will wait even after the go routine countWithChannel ends
			Go detects this problem in runtime and throws the deadlock
			To solve this we close the channel - THIS SHOULD BE DONE BY SENDER NOT RECEIVER
		*/
		c := make(chan string)
		go countWithChannelClose("sheep", c)

		for msg := range c {
			fmt.Println(msg)
		}
	case "10":
		fmt.Println("Starting: " + OPTION10)
		/*
			Here, there is a deadlock as the sender sends the message and does a blocking call waiting for a receiver
			We can resolve this by either sending the message to channel by seperate go routine or using buffered channels
		*/
		c := make(chan string)
		c <- "Hello, World"
		msg := <-c
		fmt.Println(msg)
	case "11":
		fmt.Println("Starting: " + OPTION11)
		/*
			Buffered channels are the channels with capacity
			Here, the sender channel won't block until the capacity is full
		*/
		c := make(chan string, 2)
		c <- "Hello"
		c <- "World"

		msg := <-c
		fmt.Println(msg)

		msg = <-c
		fmt.Println(msg)
	case "12":
		fmt.Println("Starting: " + OPTION12)
		/*
			Buffered channels are the channels with capacity
			Here, the sender channel won't block until the capacity is full
			If the channel is full then there will be a deadlock
		*/
		c := make(chan string, 2)
		c <- "Hello"
		c <- "World"
		c <- "Block"
		msg := <-c
		fmt.Println(msg)

		msg = <-c
		fmt.Println(msg)
	case "13":
		fmt.Println("Starting: " + OPTION13)
		/*
			Here, there are two go routines, one which is ready to send the message quickly (every 200ms), other which is slow (every 500ms)
			Eventhough, there are independent, the messages will still appear synchronous as the receiving channels are blocked as they are placed one after another
			The slower receiving channel is placed before faster receiving channel
		*/
		c1 := make(chan string)
		c2 := make(chan string)

		go func() {
			for {
				c1 <- "Every 500ms"
				time.Sleep(time.Millisecond * 500)
			}
		}()

		go func() {
			for {
				c2 <- "Every 2000ms"
				time.Sleep(time.Millisecond * 2000)
			}
		}()

		for {
			fmt.Println(<-c1)
			fmt.Println(<-c2)
		}
	case "14":
		fmt.Println("Starting: " + OPTION14)
		/*
			To resolve the synchronous receiving channels, we use select statement
			Selet statement selects the channel which is ready to send the message and prints the same
		*/
		c1 := make(chan string)
		c2 := make(chan string)

		go func() {
			for {
				c1 <- "Every 500ms"
				time.Sleep(time.Millisecond * 500)
			}
		}()

		go func() {
			for {
				c2 <- "Every 2000ms"
				time.Sleep(time.Millisecond * 2000)
			}
		}()

		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			}
		}
	case "15":
		fmt.Println("Starting: " + OPTION15)
		/*
			Queue of work to be done and multiple concurrent workers pulling items out of the queue
		*/
		jobs := make(chan int, 100)
		results := make(chan int, 100)

		// Add multiple workers
		go worker(jobs, results)
		go worker(jobs, results)
		go worker(jobs, results)
		go worker(jobs, results)
		go worker(jobs, results)

		for i := 0; i < 100; i++ {
			jobs <- i
		}
		close(jobs)

		for j := 0; j < 100; j++ {
			fmt.Println(<-results)
		}
	}

}

// Worker
// Here we are specifying two channels
// Instead of specifying bi-directional channel we specify that channel 1 will only receive and channel 2 will only send
// This will add an enforcement, if we ever try to send in channel 1, this will result in compile time error
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func countWithChannelClose(item string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- item
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}

func countWithChannel(item string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- item
		time.Sleep(time.Millisecond * 500)
	}
}

func count(item string) {
	for i := 1; true; i++ {
		fmt.Println(i, item)
		time.Sleep(time.Millisecond * 500)
	}
}

func countUptoN(item string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(item)
	}
}
