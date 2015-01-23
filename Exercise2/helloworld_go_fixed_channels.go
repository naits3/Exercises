package main

import (
    . "fmt"     // Using '.' to avoid prefixing functions with their package names
                //   This is probably not a good idea for large projects...
    "runtime"
)

var i int = 0;  

func goroutine_1(c chan int, done chan bool) {
    for n := 0;n<1000001;n++{
		i = <- c
        i++
        c <- i
	}
    done <- true
}

func goroutine_2(c chan int, done chan bool) {
    for f := 0;f<1000000;f++{
		i = <- c
        i--
        c <- i
	}
    done <- true
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())   
    
    c := make(chan int,1)
    done := make(chan bool)
    
    c <- 0
    go goroutine_1(c, done)               
	go goroutine_2(c, done)

    <- done
    <- done
    Println(<-c)
}