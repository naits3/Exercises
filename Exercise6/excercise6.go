package main

import(
		"os/exec" 
		"fmt"
		"time"
		"net"
		"strconv"
			      )

func main(){
	nr := backup()
	primary(nr)
}

func backup() (int){
	i := 0
	chListener := make(chan int)
	addr := net.UDPAddr{
		Port: 20020,
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Error listening to UDP: ",err)
		return 0
	}

	timer := time.NewTimer(time.Second)
	var buffer []byte = make([]byte, 1500)
	
	
	go func(chListener chan int){
		for {
			time.Sleep(100 * time.Millisecond)
			n, address, err := conn.ReadFromUDP(buffer)
			
			if err != nil {
				fmt.Println("Error reading from UDP: ",err)
				return
			}

			if address != nil {
				messageInt , _ := strconv.Atoi(string(buffer[0:n]))
				chListener <- messageInt
			}
		}
	}(chListener)

	for{
		select{
			case i = <-chListener:
				timer = time.NewTimer(time.Second)
			case <-timer.C:
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
		time.Sleep(500*time.Millisecond)
	}
}

func udpSender(nr_memory chan int){
	host := "127.0.0.1:20020" // IP you want to send TO. Is the port destination or source?
	conn, err := net.Dial("udp",host)

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
