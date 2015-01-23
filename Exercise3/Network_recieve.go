// Go 1.4.1

package main

import(
	"fmt"
	"net"
	"time"
)

var b []byte

func main(){
	addr := net.UDPAddr{
		Port: 30000,
		IP: net.ParseIP("129.241.187.255"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	
	if err != nil{
		fmt.Println("Error listening!")
	}


	defer conn.Close()
	for {
		time.Sleep(50 * time.Millisecond)
		n, address, err := conn.ReadFromUDP(b)
		if err != nil{
			fmt.Println("Error reading!")
		}

		if address != nil {
			fmt.Println("Got a message from",address, "with n = ", n)

			if n > 0 {
				fmt.Println("From address, ", address,", got message: ", string(b[0:n]), n)
			}
		}
	}

}

