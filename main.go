package main

import (
	"fmt"
	"log"
	"server"
	"time"
)

// implementation of mining
// 1. connect to a server
// 2. [wait mode], wait for signal 2, info about unit, or what to do
// 3. start independent loop to listen to server, for change unit info or shutdown
// 4. main loop is unit operation, sending info req by server to operate and manage unit
type connection struct {
	client *server.Client
}

var res server.Resources
var conn connection

func main() {
	res.Materials = map[string]int64{}
	fmt.Println("Client Started")

	conn.client = server.NewClient()
	go conn.Mining()
	conn.ServerConn()
}
func (c *connection) Mining() {
	var err error
	sleep := time.Second

	err = c.client.Send.Encode(&server.MsgFormat{MsgCode: 2, CInfo: c.client.AboutClient})
	if err != nil {
		log.Printf("Error in sending Info about client]] %v", err)
	}
	// starting main loop
	for ; c.client.AboutClient.Running; time.Sleep(sleep) {
		err = c.client.Send.Encode(drops())
		if err != nil {
			log.Printf("Can't sent msg: %v\n", err)
			c.client.Conn.Close()
			c.client.AboutClient.Running = false
			return
		}
	}
}
func drops() *server.MsgFormat {
	res.Materials = map[string]int64{}
	res.Materials["Rock"] += 1
	return &server.MsgFormat{MsgCode: 6, Resources: res}
}
func (c *connection) ServerConn() {
	var err error
	var msg server.MsgFormat
	for c.client.AboutClient.Running {
		msg = server.MsgFormat{}
		if err = c.client.Receive.Decode(&msg); err != nil {
			log.Printf("[Error in receiving msg]: %v", err)
			c.client.Conn.Close()
			c.client.AboutClient.Running = false
			return
		}
		switch msg.MsgCode {
		case 1: // ping, saying that server is still alive
		case 2: // get info about the server
		case 3: // signal to change settings to...
		case 4: // signal to shutdown
			log.Fatalln("Signal to shutdown at")
			c.client.Conn.Close()
			c.client.AboutClient.Running = false
		case 5: // signal to reload
		default:
			log.Println("0, something wrong")
		}
	}
}
