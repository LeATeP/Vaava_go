package main

import (
	"fmt"
	"time"
	"vaava/server"
)

func main() {
	s := server.NewServer()
	go s.AcceptConn()

	fmt.Println("server Ready")

	for {
		for i, v := range s.ClientConn {
			fmt.Println(i, v.ReceiveMsg())
		}
		time.Sleep(time.Second)
	}
}
