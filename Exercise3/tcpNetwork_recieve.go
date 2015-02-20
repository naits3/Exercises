package main

import (
        "fmt"
        "net"
)

func main() {
        service := "127.0.0.1:8080"
        listener, err := net.Listen("tcp", ":8080")
        checkError(err)
        for {
                conn, err := listener.Accept()
                //fmt.Println("Server listerning")
                _, err = conn.Read([]byte("HEAD"))
                if err != nil {
                        conn.Close()
                }
                if err != nil {
                        continue
                }
        }
}

func checkError(err error) {
        if err != nil {
                fmt.Println("Fatal error ", err.Error())
        }
}