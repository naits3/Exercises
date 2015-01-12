// Go 1.2
// go run helloworld_go.go

package main

import (
    . "fmt"     // Using '.' to avoid prefixing functions with their package names
                //   This is probably not a good idea for large projects...
    "runtime"
    "time"
)

var i int = 0;  

func goroutine_1() {
    for n := 0;n<1000000;n++{
		i++;
	}
}

func goroutine_2() {
    for n := 0;n<1000000;n++{
		i--;
	}
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())          
    go goroutine_1() 
	time.Sleep(100*time.Millisecond)                     
	go goroutine_2()
    // We have no way to wait for the completion of a goroutine (without additional syncronization of some sort)
    time.Sleep(100*time.Millisecond)
    Println(i)
}
