// Go 1.2
// go run helloworld_go.go
package main

import (
    . "fmt"     // Using '.' to avoid prefixing functions with their package names
                //   This is probably not a good idea for large projects...
    "runtime"
    "time"
    "sync"
)

var i int = 0
var mutex = &sync.Mutex{}

func goroutine_1(ch chan int) {
    for j:=0;j<1000001;j++ {
        i = <-ch
        ch <- i++
    }
}

func goroutine_2(ch chan int) {
    for n:=0;n<1000000;n++ {
        i = <-ch
        ch <- i--
    }
}

func main() {
    
    runtime.GOMAXPROCS(runtime.NumCPU())
    ch :=make(chan int)
    ch <- i
    go goroutine_1(ch)                    
    go goroutine_2(ch)
    // We have no way to wait for the completion of a goroutine (without additional syncronization of some sort)
    time.Sleep(100*time.Millisecond)
    Println(i)
}
