package main

import(
		"os/exec" 
		"fmt"
		"strconv"
		"time"
			      )

func main(){
	nr := backup()
	primary(nr)
}

func backup(){

	chListener = make(chan int)
	addr := net.UDPAddr{
		Port: 20020,
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil s{
		fmt.Println("Error listening to UDP: ",err)
		return
	}

	timer := time.NewTimer(time.Second)
	var buffer []byte = make([]byte, 1500)
	defer conn.Close()
	
	go func listen(chListener chan bool){
		for {
			time.Sleep(100 * time.Millisecond)
			n, address, err := conn.ReadFromUDP(buffer)
			if err != nil {
				fmt.Println("Error reading from UDP: ",err)
				return
			}

			if address != nil {
				chListener <- int(buffer[0:n])
			}
		}
	}

	for{
		select{
			case i := <-chListener:
				timer = time.NewTimer(time.Second)
			case <-timer:
				return i
		}
	}
}

func primary(int nr){
	createNewBackup();
	nr_memory := make(chan int)
	go udpSender(nr_memory);
	for i:= nr; i++{
		nr_memory <- i
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
		_ , err := conn.Write([]byte(i))
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