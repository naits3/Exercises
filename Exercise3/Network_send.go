// Go 1.4.1

package main

import(
	"fmt"
	"net"
	"time"
)



func main(){
	addr := net.UDPAddr{
		Port: 20023,
		IP: net.ParseIP("129.241.187.255"),
	}

	byteArray := []byte("h")
	fmt.Println(byteArray)

	conn, err2 := net.DialUDP("udp",nil, &addr)
	if err2 != nil{
			fmt.Println("Error when conn")
		}

	for{

		n,err1 := conn.Write(byteArray)
		if err1 != nil{
			fmt.Println(err1)
		}
		if(n>0){
			fmt.Println("Yes")
		}
		time.Sleep(1000* time.Millisecond)
	}
}