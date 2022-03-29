package main

import (
	"fmt"
	"log"
	"time"
	"vaava/psql"
	"vaava/server"
)

func main() {
	s := server.NewClient()
	res := server.Resources{map[string]uint64{}}
	sleep := time.Second / 10

	for ;s.AboutClient.Running ;time.Sleep(sleep) {
		res.Materials["Ore"] += 1
		if res.Materials["Ore"] > 10 {
			err := s.Send.Encode(&server.MsgFormat{
				MsgCode: 6, 
				Resources: res})
			if err != nil {
				log.Printf("Can't sent msg: %v\n", err)
				s.Conn.Close()
				s.AboutClient.Running = false
				break
			}
			res = server.Resources{map[string]uint64{}}
		}
	}
}
func ServerConn(s *server.Client) {
	sleep := time.Second
	for ;;time.Sleep(sleep) {
		msg := &server.MsgFormat{}
		err :=s.Receive.Decode(msg)
		if err != nil {
			log.Printf("[Error in receiving msg]: %v", err)
			s.Conn.Close()
			s.AboutClient.Running = false
		}
		switch msg.MsgCode {
		case 1: // ping, saying that server is still alive
		case 2: // get info about the server
		case 3: // signal to change settings to...
		case 4: // signal to shutdown
			log.Printf("Signal to shutdown at %v\n", time.Now().UTC())
		case 5: // signal to reload
		default:
			log.Println("0, something wrong")
		}
	}
}
















func query(db psql.DbInterface, cmd string) error {
    data, err := db.QuerySelect("select * from items;")
    if err != nil { log.Fatalln(err) }

    for _, w := range data {
		fmt.Printf("[%v] %-15v %d\n", w["id"], w["name"], w["amount"])
	}
	return nil
}