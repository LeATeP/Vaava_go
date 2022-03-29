package main

import (
	"fmt"
	"log"
	"time"
	"vaava/server"
)

func main() {
	s 	:= server.NewServer()
	go s.AcceptConn()

	fmt.Println("server ready")

	res := server.Resources{map[string]uint64{}}
	go receivingLoop(s, &res)
	time.Sleep(time.Second * 100)

	// sql handler, update and send clients everything that is needed

}

func receivingLoop(s *server.Server, res *server.Resources) {
	sleep := time.Second / 10
	for ;;time.Sleep(sleep) {
		for i, v := range s.ClientConn {
			msg  := server.MsgFormat{}
			err  := v.Receive.Decode(&msg)
			if err != nil {
				log.Printf("%v [err]: %v\n", i, err)     // well would be to put client identifiers like containerId and stuff
				v.Conn.Close()
				delete(s.ClientConn, i)
				continue
			}
			switch msg.MsgCode {
			case 1: // get ping that client is active
			case 2: // get info about client
				v.AboutClient = &msg.CInfo
			case 3: // something changed in client
			case 4: // client shutting down
				v.Conn.Close()
				delete(s.ClientConn, i)
			case 5: // client reloading
			case 6: // update resources
				for i, k := range msg.Resources.Materials {
					res.Materials[i] += k
					fmt.Println(res.Materials)
				}
			default:
				fmt.Println("0, something wrong")
			}
		}
	}
}