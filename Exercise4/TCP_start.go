
package main

import (
	"fmt"
	"net"
	"time"
)

var connectionList []string

func ReceivePack(port string, receive chan []byte) []byte {
	listener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		return []byte("0")
	}

	for {
		var buffer []byte = make([]byte, 1024)
		conn, err := listener.Accept()
		packetSize, err := conn.Read(buffer)

		if err != nil {
			return []byte()
		}

		if packetSize > 0{
			receive <- buffer[0:n]
		}
	}
}

func SendPack(pack []byte, host string, chSend chan bool){
	addr, _ := net.ResolveTCPAddr("tcp",host)
	conn, err := net.DialTCP("tcp", nil, addr)

	if err != nil{
		return
	}
	
	_, err := conn.Write(pack)
	
	if err != nil {
		return
	}

	chSend <- true
}


func listenForConnections(port string){
	addr := net.UDPAddr{
		Port: port
	}

	conn, err := net.ListenUDP("udp", &addr)

	if err != nil {
		fmt.Println("Error listening to UDP: ",err)
		return
	}

	var buffer []byte = make([]byte, 1024)
	
	defer conn.Close()
	for {
		_, address, err := conn.ReadFromUDP(buffer)

		if err != nil {
			fmt.Println("Error reading from UDP: ",err)
			return
		}

		if address != nil {
			append(connectionList, address.IP)
			// SEND OK HERE and length of ConnectionList.
		}
	}
}

func requestConnection(ip string, port string, )

func main(){
	port := "20020"
	host := "78.91.38.8"+":"+port

	chReceive := make(chan []byte)
	chSend := make(chan bool)	
	
	packetBuffer := make([]byte, 1024)

	go ReceivePack(port, chReceive)
	go SendPack([]byte("Hei, Stian"), host, chSend)

	for {
		select {
			case m = <- chReceive:
				fmt.Printf("Received: %s\n", m)
				go SendPack([]byte("Hei, Stian"), host, chSend)
			case <- chSend:
				fmt.Printf("sent!")
		}
	}
}