package main

import (
	"four/four/cmd"
	"log"
)

func main(){
	err:= cmd.Execute()
	if err != nil {
		log.Fatal("cmd.Execute err:", err)
	}
}
