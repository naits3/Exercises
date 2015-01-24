//Go 1.4.1

package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	host := "129.241.187.136:80"
	addr, _ := net.ResolveTCPAddr("tcp",host)
	
	conn, err := net.DialTCP("tcp", nil, addr)

	if err != nil{
		fmt.Println("Error connecting to ", host, ": ",err)
		return
	}

	fmt.Println("Connected to server at ", host)

	defer conn.Close()

	fmt.Println("About to write to connection..")

	for {
		time.Sleep(1000 * time.Millisecond)
		n, err := conn.Write([]byte("Connect to: 129.241.187.158:20010\\0"))
		if err != nil {
			fmt.Println("Error writing to server: ", err)
			return
		}

		if n > 0 {
			fmt.Println("Wrote ",n," bytes to server at ", host)
		}

	}
}