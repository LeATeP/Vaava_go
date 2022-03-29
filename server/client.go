package server

import (
	"encoding/gob"
	"log"
	"net"
	"time"
)

func NewClient() *Client {
	conn, err := net.Dial("tcp", connectToLocal) // listen for clients
	if err != nil {
		log.Printf("[failed to connect]: %v\n", err)
	}
	return &Client{
		Conn: 		 	  conn,
		Receive: 	 	  gob.NewDecoder(conn),
		Send: 			  gob.NewEncoder(conn),
		AboutClient: &AboutClientInfo{
			Id: 		  1,
			Name: 		  "mining",
			Status: 	  "mine",
			Start: 		  time.Now().UTC(),
			Running:  	  true,
			ContainerId:  "693210",
			TickDataSend: time.Second,
		},
	}
}
