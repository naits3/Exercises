// Go 1.4.1

package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	addr := net.TCPAddr{
		Port: 8,
		IP: net.ParseIP("129.241.187.255"), //Why does this IP have to be BroadCast?
	}

	// Create a socket:
	fmt.Println("Creating a socket..")
	listener, err := net.ListenTCP("tcp", &addr)

	
	if err != nil {
		fmt.Println("Error listening to TCP: ",err)
		return
	}

	conn, err := listener.Accept()

	fmt.Println("..Socket created! Listening to Port ",addr.Port)

	var buffer []byte = make([]byte, 1500)
	defer conn.Close()

	for {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("h")

		n, err := conn.Read(buffer)

		if err != nil {
			fmt.Println("Error reading from TCP: ",err)
			return
		}


		fmt.Println("Got message with n = ", n)

		if n > 0{
			fmt.Println("got message: ", string(buffer[0:n]), n)
		}
		

	}
}