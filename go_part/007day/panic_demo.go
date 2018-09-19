// blocking.go
// throw: all goroutines are asleep - deadlock!
package main

import (
	"fmt"
	"time"
)

func main() {

	var c1 chan string = make(chan string)

	go func() {

		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	fmt.Println("c1 is ", <-c1)

}
