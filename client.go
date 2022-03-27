package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"time"
)
type connection struct {
	sleepTime    time.Duration
	connection   net.Conn
	sendConn 	*gob.Encoder
	receiveConn *gob.Decoder
}
type serverHandler interface {
	init_ConfigHandler(net.Conn)
	server_Loop()
	sendMsg(string) error
	closeConn()
}

func client() (serverHandler, error) {
	conn, err := net.Dial("tcp", ":9000")
	if err != nil { return nil, err }

	var serverHandler serverHandler = &connection{}
	serverHandler.init_ConfigHandler(conn)
	

	go serverHandler.server_Loop()
	return serverHandler, nil
}

func (c *connection) init_ConfigHandler(conn net.Conn) {
	c.sleepTime   = time.Second / 10
	c.connection  = conn
	c.receiveConn = gob.NewDecoder(c.connection)
	c.sendConn 	  = gob.NewEncoder(c.connection)

}
func (c connection) server_Loop() {
	for {
		var msg string
		err := c.receiveConn.Decode(&msg)
		if err != nil {
			log.Printf("conn ended: %v\n", err)
			break
		}
		fmt.Printf("msg: %v\n", msg)
		time.Sleep(c.sleepTime)
	}
	c.connection.Close()

}
func (c connection) sendMsg(msg string) error {
	err := c.sendConn.Encode(msg)
	if err != nil {
		return err }
	return nil
}
func (c connection) closeConn() {
	c.connection.Close()
}