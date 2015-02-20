package main

import(
		"os/exec" 
		"fmt" 
		"io"    )


func main(){
	CreateBackup()
	 for i := 0; i < 100; i++{
        fmt.Println(i)
    }
}


func primary(){

}

func backup(){

}

func CreateBackup(){
	cmd := exec.Command("cmd","/C","start")
	err := cmd.Start()
	
	if err != nil {
    	fmt.Println(err)
    }
   writer, err := exec.Stdin(&cmd)
   writer. 


}