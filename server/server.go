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
type clientHandler interface {
	init_ConfigHandler(net.Conn)
	client_Loop()
	sendMsg(string) error
	closeConn()
}
func main() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil { log.Fatalln(err.Error()) }
	fmt.Println("server is up")
	
	for {
		var userHandler clientHandler = &connection{}

		conn, err := ln.Accept() // listen for clients
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("connected: %v\n", conn)
		userHandler.init_ConfigHandler(conn)
		go userHandler.client_Loop()

	}
}
func (c *connection) init_ConfigHandler(conn net.Conn) {
	c.sleepTime   = time.Second / 10
	c.connection  = conn
	c.receiveConn = gob.NewDecoder(c.connection)
	c.sendConn 	  = gob.NewEncoder(c.connection)

}
func (c connection) client_Loop() {
	for {
		var msg string
		err := c.receiveConn.Decode(&msg)
		if err != nil {
			log.Printf("conn ended: %v\n", err)
			break
		}
		fmt.Printf("msg: %v", msg)
		time.Sleep(c.sleepTime)
	}
	c.connection.Close()

}
func (c connection) sendMsg(msg string) error {
	err := c.sendConn.Encode(msg)
	if err != nil {
		log.Printf("failed to send msg: %v\n", err)
		return err
	}
	return nil
}
func (c connection) closeConn() {
	c.connection.Close()
}