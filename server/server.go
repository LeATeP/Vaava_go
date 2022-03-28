package server

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"time"
)
func NewServer() *Server {
	info   := Info{
		Id: "0",
		Name: "main",
		ContainerId: "",
		MaxLoad: 100,
	}
	ls, _  := net.Listen("tcp", listenToLocal)
	server := Server{
		Start: time.Now().UTC(),
		Info: info,
		Running: true,
		Stats: &Stats{},
		ClientConn: map[uint64]*Client{},
		Listener: ls,
	}
	return &server
}
func (s *Server) AcceptConn() {
	var i uint64
	for ; s.Running ;i++ {
		conn, err := s.Listener.Accept() // listen for clients
		if err != nil {
			log.Printf("[failed to connect]: %v\n", err)
		}
		fmt.Printf("connected [%v]: %v\n", i, conn)

		s.ClientConn[i] = &Client{
			Conn: conn, 
			Receive: gob.NewDecoder(conn), 
			Send: gob.NewEncoder(conn)}
		}
}
func (c Client) ReceiveMsg() MsgFormat {
	msg := &MsgFormat{}
	c.Receive.Decode(msg)
	return *msg
}
func (c Client) SendMsg(msg *MsgFormat) error {
	return c.Send.Encode(msg)
}
func (c Client) CloseConn() {
	c.Conn.Close()
}