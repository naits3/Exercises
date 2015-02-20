package main

import(
		"os/exec" 
		"fmt"
		"time"
		"net"
		"strconv"
			      )

func main(){

	host := "localhost:20020"
	udpAdr, _ := net.ResolveUDPAddr("udp",host)
	conn, _ := net.ListenUDP("udp",udpAdr)

	nr := backup(conn)
	conn.Close()

	primary(nr)
	
}

func backup(conn net.Conn) (int){
	i := 0
	chListener := make(chan int)
	
	go func(chListener chan int, conn net.Conn){
		var buffer []byte = make([]byte, 1500)		
				
		for {
			n, _ := conn.Read(buffer)
			messageInt , _ := strconv.Atoi(string(buffer[0:n]))
			chListener <- messageInt
		}
		
	}(chListener, conn)
	
	for{
		select{
			case i = <- chListener:
				continue
			case <-time.After(time.Second):
				return i
		}
	}
}


func primary(nr int){
	createNewBackup();
	nr_memory := make(chan int)
	
	go udpSender(nr_memory)
	i := nr
	for {
		nr_memory <- i
		fmt.Println(i)
		i++
		time.Sleep(100*time.Millisecond)
	}
}

func udpSender(nr_memory chan int){
	host := "localhost:20020" // IP you want to send TO. Is the port destination or source?
	udpAdr, _ := net.ResolveUDPAddr("udp",host)
	conn, err := net.DialUDP("udp",nil,udpAdr)

	if err != nil{
		fmt.Println("Error connecting to ", host, ": ",err)
		return
	}

	defer conn.Close()
	
	for {
		i := <- nr_memory 
		s := strconv.Itoa(i)
		_ , err := conn.Write([]byte(s))
		if err != nil {
			fmt.Println("Error writing to server: ", err)
			return
		}

	}
}


func createNewBackup(){
	cmd := exec.Command("gnome-terminal","-x","sh", "-c", "go run excercise6.go")
	err := cmd.Start()
	if err != nil {
    	fmt.Println(err)
    }
}
