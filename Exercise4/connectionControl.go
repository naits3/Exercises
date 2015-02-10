package network

import (
	"fmt"
	"net"
)

var connMap map[string]net.Conn = nil
var port string = nil							//Decide a listnening port!!


// conn is connection

func InitConnControl(){
	//Run connectionConntroll thread
	//Run requestConnections once to see if there are other connections ready
	//Starts listenForConnection from port
}

func runConnControl(){
	// Select for Connection controll legges her
	// Tar inn alle nødvendige channels for kommunikasjon
}

func requestConn(ip string, port string, ){
	//UDP broadcast en gang ved oppstart
	//hvis ikke connection prøv igjen altså hvis det kommer en error
}

func listenForNewConn(port int){
	addr := net.UDPAddr{
		Port: port,
	}

	conn, err := net.ListenUDP("udp", &addr)

	if err != nil {
		fmt.Println("Error listening to UDP: ",err)
		return
	}

	var buffer []byte = make([]byte, 1024)
	
	defer conn.Close()
	for {
		n, address, err := conn.ReadFromUDP(buffer)

		if err != nil {
			fmt.Println("Error reading from UDP: ",err)
			return
		}

		if address != nil && string(buffer[0:n])=="CONNECTME" {
			// SEND OK HERE
		}
	}
}


func addConn(ip string){
	//Legger til ip som hash variabel med en connection TCP
}

func removeConn(ip string){
	//Fjerner en connection fra dictonary for TCP
}

func GetConnMap() map[string]net.Conn{
	// Sender alle TCP til kommunikasjonskontroll
	// Må ha variable kontroll med addConn.. og deleteConn..
}

// Ikke ferdig!!!
// Sette en egen variabel for broadcast signal??
func getBroadcastConn() string{
	ifaces, _ := net.Interfaces()
	// handle err
	for _, i := range ifaces {
	    addrs, _ := i.Addrs()
	    // handle err
	    for _, addr := range addrs {
	        switch v:= addr.(type) {
	        case *net.IPAddr:
	            if(v.IP.String() != "0.0.0.0"){
	            	return(v.IP.String())
	        	}
	        }
	    }
	}
	return "is_offline"
}

func sendPingToConn(){
	//UDP broadcast that listens for broken connections in the map
	//For each signal sent the number of replies should be equal to number of connections
	//Var tidligere ping funksjonen
}

func recievePingFromConn(){
	//Keeps track on connections avaliability
}



