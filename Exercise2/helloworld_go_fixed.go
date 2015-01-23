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

func goroutine_1() {
    for j:=0;j<1000001;j++ {
        mutex.Lock()
        i++
        mutex.Unlock()
    }
}

func goroutine_2() {
    for n:=0;n<1000001;n++ {
        mutex.Lock()
        i--
        mutex.Unlock()
    }
}

func main() {
    
    runtime.GOMAXPROCS(runtime.NumCPU())
    go goroutine_1()                    
    go goroutine_2()
    // We have no way to wait for the completion of a goroutine (without additional syncronization of some sort)
    time.Sleep(400*time.Millisecond)
    Println(i)
}
