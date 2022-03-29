package main

import (
	"fmt"
	"time"
	"vaava/server"
)

func main() {
	s := server.NewServer()
	go s.AcceptConn()

	fmt.Println("server ready")

	for {
		for i, v := range s.ClientConn {
			msg := server.MsgFormat{}
			err := v.Receive.Decode(&msg)
			fmt.Println(msg, err)
		}
		time.Sleep(time.Second)
	}
}
