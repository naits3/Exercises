
package main

import "time"

func backup(){
	addr := net.UDPAddr{
		Port: 127.0.0.1,
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil s{
		fmt.Println("Error listening to UDP: ",err)
		return
	}

	timer := time.NewTimer(time.Second)
	var buffer []byte = make([]byte, 1500)
	defer conn.Close()

	go func listen(){
		for {
			time.Sleep(100 * time.Millisecond)
			n, address, err := conn.ReadFromUDP(buffer)
			if err != nil {
				fmt.Println("Error reading from UDP: ",err)
				return
			}

			if address != nil {
				// MOTTAR MELDING HER
				}
			}
		}
	}
}