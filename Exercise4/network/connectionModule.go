package network

import (
	"fmt"
	"net"
	"time"

)

var connMap map[string]net.Conn = nil
var port string = nil							//Decide a listnening port!!

// conn is connection

func InitConnModule(){
	//init all channels
	//Create broadcast connection
	//Thread connDurationHandler()
	//Thread listenPing()
	//Thread sendPing()
}

func addConn(ip string){
	//adds a connection to map with ip as key
}

func removeConn(ip string){
	//remove conn from map with ip string
}

func GetConnMap() map[string]net.Conn{
	// Sender alle TCP til kommunikasjonskontroll
	// Må ha variable kontroll med addConn.. og deleteConn..
}

func sendPing(udpConn net.Conn){
	// Ping all in subnet every 100ms on broadcast
}

func listenPing(port string, chListenPing chan){
	// listning for broadcast signals
	// if valid signal send to connDurationHandler
}

func connDurationHandler(){
	// start new timer for 1 sec
	// Get all keys
		for{select
			case: time:=<-Timer
				//check if the list are empty, if not remove ip from map
				//start new timer for 1 sec
			case: //gets ip string from listenPing() channel
				//check if it is in map
				//if in map remove from list
				//if not in map add new connection to map

		}
}

func createBroadcastConn() net.UDPConn{
	//getHostIP()
	//Use 255 as last part in IP
	//make UDP connection
}

func GetPort() string{
	return port
}

func getHostIP() string{
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



