//Go 1.4.1

package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	host := "129.241.187.136:20010" // IP you want to send TO. Is the port destination or source?
	conn, err := net.Dial("tcp",host)

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