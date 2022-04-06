package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"server"
)
type connection struct {
	server *server.Server
}
var i int64  				// client connection number
var res server.Resources	// data client sending to process
var pool connection			
func main() {
	res.Materials = map[string]int64{}
	pool.server   = server.NewServer()
	fmt.Println("server ready")

	for ;;i++{
		pool.AcceptConn(i)
	}
}
func (c *connection) AcceptConn(i int64) {
	conn, err := c.server.Listener.Accept() // listen for clients
	if err != nil {
		log.Printf("[failed to connect]: %v\n", err)
	}
	fmt.Printf("connected [%v]: %v\n", i, conn)

	c.server.ClientConn[i] = &server.Client{
		Conn: 	 conn, 
		Receive: gob.NewDecoder(conn), 
		Send: 	 gob.NewEncoder(conn),
	}
	go c.ManageConnection(i)

}

func (c *connection) ManageConnection(i int64) {
	var count int64 
	var err   error
	var msg   server.MsgFormat
	client    := c.server.ClientConn[i]
	for {
		msg    = server.MsgFormat{}
		if err = client.Receive.Decode(&msg); err != nil {
			log.Printf("%v [err]: %v\n", i, err)     // well would be to put client identifiers like containerId and stuff
			client.Conn.Close()
			delete(c.server.ClientConn, i)
			return
		}
		switch msg.MsgCode {
		case 1: // get ping that client is active
		case 2: // get info about client
			client.AboutClient = &msg.CInfo
		case 3: // something changed in client
		case 4: // client shutting down
			client.Conn.Close()
			delete(c.server.ClientConn, i)
			return
		case 5: // client reloading
		case 6: // update resources
			for i, k := range msg.Resources.Materials {
				res.Materials[i] += k
				fmt.Println(res.Materials)
			}
		
			if count == 2 {
				client.Send.Encode(server.MsgFormat{MsgCode: 4})
				client.Conn.Close()
				delete(c.server.ClientConn, i)
				return
			}
			count++
		default:
			fmt.Println("0, something wrong")
		}
	}
}

// c.Send.Encode
// c.Receive.Decode
// c.Conn.Close 
// work just fine, but if needed to be put in interface, then it will be necessary? 
