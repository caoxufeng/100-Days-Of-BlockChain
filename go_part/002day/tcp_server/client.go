package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	//open conn
	conn, err := net.Dial("tcp", "localhost:50007")

	if err != nil {
		//can not connect because the target compute refuse.
		fmt.Println("Error dialing.", err.Error())
		return
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First,what is your name?")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\r\n")

	//send the message to the server util system exit.
	for {

		fmt.Println("What is send to the server? Type Q to quit.")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return
		}

		_, err = conn.Write([]byte(trimmedClient + " says:" + trimmedInput))

	}

}
