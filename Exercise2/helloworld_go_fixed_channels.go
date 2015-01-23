package main

import (
    . "fmt"     // Using '.' to avoid prefixing functions with their package names
                //   This is probably not a good idea for large projects...
    "runtime"
    
)

var i int = 0; 
var c chan int 

func goroutine_1() {
    for n := 0;n<1000001;n++{
		i := <- c
        c <- (i+1)
	}
}

func goroutine_2() {
    for n := 0;n<1000000;n++{
		i := <- c
        c <- (i-1)
	}
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())   
    
    c := make(chan int)
    c <- 0
    
    go goroutine_1()               
	go goroutine_2()

    Println(<-c)
}