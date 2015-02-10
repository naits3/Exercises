package network

import (
	"fmt"
	"net"
)

var connMap map[string]net.Conn = nil
var port string = nil							//Decide a listnening port!!


// conn is connection

func InitConnModule(){
	//init all channls
	//Run connHandler thread
	//Run requestConnections once to see if there are other connections ready
	//Starts listenForConnection from port
	//Thread receivePingFromConn
	//Thread sendPingToConn
}

func runConnHandler(){
	// Select for Connection controll legges her
	// Tar inn alle nødvendige channels for kommunikasjon
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
func getBroadcastConn() string err{
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



