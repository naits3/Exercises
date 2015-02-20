
package main

import "time"

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
