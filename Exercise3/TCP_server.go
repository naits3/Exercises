// Go 1.4.1

package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	

	// Create a socket:
	fmt.Println("Creating a socket..")
	listener, err := net.Listen("tcp", ":20020")

	
	if err != nil {
		fmt.Println("Error listening to TCP: ",err)
		return
	}

	

	fmt.Println("..Socket created! Listening to Port ")

	var buffer []byte = make([]byte, 1024)
	//defer conn.Close()

	for {

		conn, err := listener.Accept()

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