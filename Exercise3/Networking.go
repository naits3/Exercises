// Go 1.4.1

package main

import(
	"fmt"
	"net"
)

func main(){
	addr := net.UDPAddr{
		Port: 30000,
		IP: net.ParseIP("127.0.0.1"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	
	if err != nil{
		fmt.Println("Error")
	}

	
}