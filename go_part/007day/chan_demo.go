package main

import "fmt"
import "time"

func main() {
	c := make(chan int)
	go func() {
		time.Sleep(15 * 1e9)
		x := <-c
		fmt.Println("received", x)
	}()
	fmt.Println("sending", 10)
	c <- 10 //因为channel  被阻塞，所以阻止了程序的结束
	fmt.Println("sent", 10)
}
