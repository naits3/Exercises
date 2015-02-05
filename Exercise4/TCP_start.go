
package main

import (
	"fmt"
	"net"
	"time"
	"os"
)


func receivePack(port string) string {
	fmt.Println("Creating a socket..")
	listener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		fmt.Println("Error listening to TCP: ",err)
		return
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
			return
		}

		if n > 0{
			receive <- string(buffer[0:n])
		}
		
	}
}

func sendPack(pack string, host string, done chan bool){
	//host := "78.91.36.36:80"
	addr, _ := net.ResolveTCPAddr("tcp",host)
	conn, err := net.DialTCP("tcp", nil, addr)

	if err != nil{
		fmt.Println("Error connecting to ", host, ": ",err)
		return
	}

	fmt.Println("Connected to server at ", conn.RemoteAddr().Network())
	fmt.Println("About to write to connection..")

	n, err := conn.Write([]byte(pack))
	if err != nil {
		fmt.Println("Error writing to server: ", err)
		return
	}

	if n > 0 {
		fmt.Println("Wrote ",n," bytes to server at ", host)
	}

	done <- true
}

func promptSend(){

}

func main(){

	port := "80"
	chReceive := make(chan string)
	chSend := make(chan bool)	
	string m
	
	go receivePack(port)
	go send

	for {
		select {
			case m <- chReceive:
				Println("Received: ",)
			case <- chSend:
				print("sent!")
		}
	}

}