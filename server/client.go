package server

import (
	"encoding/gob"
	"log"
	"net"
	"os"
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
			Start: 		  time.Now().UTC(),
			Running:  	  true,
			ContainerId:  getHostName(),
		},
	}
}

// return hostname from the env variable
func getHostName() string {
	return os.Getenv("HOSTNAME")
}