
package main

import (
	"fmt"
	"net"
	"time"
)


func ReceivePack(port string, receive chan []byte) []byte {
	fmt.Println("Creating a socket..")
	listener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		fmt.Println("Error listening to TCP: ",err)
		return []byte("0")
	}

	fmt.Println("..Listening..")

	var buffer []byte = make([]byte, 1024)

	for {

		conn, err := listener.Accept()
		fmt.Println("Connected to ",conn.RemoteAddr().String())

		time.Sleep(100 * time.Millisecond)

		n, err := conn.Read(buffer)

		if err != nil {
			fmt.Println("Error reading from TCP: ",err)
			return []byte("0")
		}

		if n > 0{
			receive <- buffer[0:n]
		}
		
	}
}

func SendPack(pack []byte, host string, chSend chan bool){
	//host := "78.91.36.36:80"
	addr, _ := net.ResolveTCPAddr("tcp",host)
	conn, err := net.DialTCP("tcp", nil, addr)

	if err != nil{
		fmt.Println("Error connecting to ", host, ": ",err)
		return
	}

	fmt.Println("Connected to server at ", conn.RemoteAddr().Network())
	fmt.Println("About to write to connection..")

	for {
		time.Sleep(1000 * time.Millisecond)
		
		n, err := conn.Write(pack)
		if err != nil {
			fmt.Println("Error writing to server: ", err)
			return 
		}

		if n > 0 {
			fmt.Println("Wrote ",n," bytes to server at ", host)
		}

		chSend <- true
	}
}

func main(){
	port := "20020"
	host := "78.91.38.8"+":"+port
	chReceive := make(chan []byte)
	chSend := make(chan bool)	
	
	go ReceivePack(port, chReceive)
	go SendPack([]byte("Hei, Stian"), host, chSend)

	for {
		select {
			case <- chReceive:
				fmt.Println("Received: %s", chReceive)
			case <- chSend:
				fmt.Println("sent!")
		}
	}
}