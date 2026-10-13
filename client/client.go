package main

import (
	"flag"
	"net"
	"fmt"
)

func read(conn net.Conn) {
	//TODO In a continuous loop, read a message from the server and display it.
	for {
		var msg string
		_, err := fmt.Fscanln(conn, &msg)
		if err != nil {
			fmt.Println("Error reading from server:", err)
			return
		}
		fmt.Println("Received from server:", msg)
	}
}

func write(conn net.Conn) {
	//TODO Continually get input from the user and send messages to the server.
	var input string
	for {
		fmt.Print("Enter message to send: ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		_, err = fmt.Fprintln(conn, input)
		if err != nil {
			fmt.Println("Error sending message to server:", err)
			return
		}
	}

}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()
	//TODO Try to connect to the server
	ln, err := net.Dial("tcp", *addrPtr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer ln.Close()

	//TODO Start asynchronously reading and displaying messages
	go read(ln)
	//TODO Start getting and sending user messages.
	write(ln)
}
