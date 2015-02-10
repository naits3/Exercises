package network

import (
	"fmt"
	"net"
	"main"
)

var connMap map[string]net.Conn = nil
var port string = nil							//Decide a listnening port!!


// conn is connection

func InitConnModule(){
	//init all channls
	//Thread Run connHandler
	//Create broadcast connection
	//Thread listenPing
	//Thread sendPing
}

func runConnHandler(){
	// Select for Connection controll legges her
	// takes in necessary channels
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

func sendPing(udpConn net.Conn){
	//UDP broadcast that listens for broken connections in the map
	//For each signal sent the number of replies should be equal to number of connections
	//Var tidligere ping funksjonen
}

func listenPing(port string){
	//Keeps track on connections avaliability
}

// Ikke ferdig!!!
// Sette en egen variabel for broadcast signal??
func createBroadcastConn() string{
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

func GetPort() string{
	return port
}

func getHostIP(){

}



