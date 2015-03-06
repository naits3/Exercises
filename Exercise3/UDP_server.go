//Go 1.4.1

package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	addr := net.UDPAddr{
		Port: 3000,
		IP: net.ParseIP(":255"), //Why does this IP have to be BroadCast?
	}

	// Create a socket:
	fmt.Print("Creating a socket..")
	conn, err := net.ListenUDP("udp", &addr)

	
	if err != nil {
		fmt.Println("Error listening to UDP: ",err)
		return
	}

	fmt.Println("..Socket created! \nListening to Port ",addr.Port)

	var buffer []byte = make([]byte, 1500)
	defer conn.Close()

	for {
		time.Sleep(100 * time.Millisecond)

		n, address, err := conn.ReadFromUDP(buffer)

		if err != nil {
			fmt.Println("Error reading from UDP: ",err)
			return
		}

		if address != nil {

			fmt.Println("Got message from ",address," with n = ", n)

			if n > 0{
				fmt.Println("From address", address, "got message: ", string(buffer[0:n]))
			}
		}

	}
}