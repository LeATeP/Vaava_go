package server

import (
	"encoding/gob"
	"log"
	"net"
	"os"
	"time"
)

func NewClient() *Client {
	conn, err := net.Dial("tcp", connectToLocal) // connect to a server
	if err != nil {
		log.Printf("[failed to connect]: %v\n", err)
	}
	return &Client{
		Conn: 		 	  conn,
		Receive: 	 	  gob.NewDecoder(conn),
		Send: 			  gob.NewEncoder(conn),
		AboutClient: AboutClientInfo{
			Start: 		  time.Now().UTC(),
			Running:  	  true,
			ContainerId:  os.Getenv("HOSTNAME"),
		},
	}
}
