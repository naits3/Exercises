package network

import (
	"fmt"
	"net"
)

// Comm = Communication

func initCommModule(){
	//initaliserer all channels(receive chan, chsend
	//starts communication control
	//starts also connectionControll
}

func runCommHandler(){
	//Select which controls sending and receiving
}

func listenPack(port string, chReceive chan []byte){
	listener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		return []byte("0")
	}

	for {
		var buffer []byte = make([]byte, 1024)
		conn, err := listener.Accept()
		packetSize, err := conn.Read(buffer)

		if err != nil {
			return []byte("0")
		}

		if packetSize > 0{
			receive <- buffer[0:packetSize]
		}
	}
}

func SendPack(pack []byte, host string, chSend chan bool){
	addr, _ := net.ResolveTCPAddr("tcp",host)
	conn, err := net.DialTCP("tcp", nil, addr)

	if err != nil{
		return
	}
	
	_, err = conn.Write(pack)
	
	if err != nil {
		return
	}

	chSend <- true
}


