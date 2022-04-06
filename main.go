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
	server *server.Client
}
var res  server.Resources
var conn connection
func main() {
	res.Materials = map[string]int64{}
	fmt.Println("Client Started")
	
	conn.server = server.NewClient()
	go conn.Mining()
	conn.ServerConn()
}
func (c *connection) Mining() {
	sleep := time.Second / 10

	for ;c.server.AboutClient.Running ;time.Sleep(sleep) {
		res.Materials["Ore"] += 1


		if res.Materials["Ore"] > 10 {
			err := c.server.Send.Encode(
				&server.MsgFormat{
					MsgCode: 6, 
					Resources: res})
			if err != nil {
				log.Printf("Can't sent msg: %v\n", err)
				c.server.Conn.Close()
				c.server.AboutClient.Running = false
				return
			}
			res.Materials = map[string]int64{}
		}
	}
}
func (c *connection) ServerConn() {
	var err   error
	var msg   server.MsgFormat
	for ;c.server.AboutClient.Running; {
		msg    = server.MsgFormat{}
		if err = c.server.Receive.Decode(&msg); err != nil {
			log.Printf("[Error in receiving msg]: %v", err)
			c.server.Conn.Close()
			c.server.AboutClient.Running = false
		}
		switch msg.MsgCode {
		case 1: // ping, saying that server is still alive
		case 2: // get info about the server
		case 3: // signal to change settings to...
		case 4: // signal to shutdown
			log.Fatalf("Signal to shutdown at %v\n", time.Now().UTC())
		case 5: // signal to reload
		default:
			log.Println("0, something wrong")
		}
	}
}