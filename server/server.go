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
		Id: 		 "0",
		Name: 		 "main",
		ContainerId: "",
		MaxLoad: 	 100,
	}
	ls, _  := net.Listen("tcp", listenToLocal)
	server := Server{
		Start: 		time.Now().UTC(),
		Info: 		info,
		Running: 	true,
		Stats: 		&Stats{},
		ClientConn: map[int64]*Client{},
		Listener: 	ls,
	}
	return &server
}
func (s *Server) AcceptConn() {
	var i int64
	for ; s.Running ;i++ {
		conn, err := s.Listener.Accept() // listen for clients
		if err != nil {
			log.Printf("[failed to connect]: %v\n", err)
			continue
		}
		fmt.Printf("connected [%v]: %v\n", i, conn)

		s.ClientConn[i] = &Client{
			Conn: 	 conn, 
			Receive: gob.NewDecoder(conn), 
			Send: 	 gob.NewEncoder(conn),
		}
	}
}

// c.Send.Encode
// c.Receive.Decode
// c.Conn.Close 
// work just fine, but if needed to be put in interface, then it will be necessary? 

// func (c Client) ReceiveMsg() (MsgFormat, error) {
	// msg := &MsgFormat{}
	// return *msg, c.Receive.Decode(msg)
// }
// func (c Client) SendMsg(msg *MsgFormat) error {
	// return c.Send.Encode(msg)
// }
// func (c Client) CloseConn() {
	// c.Conn.Close()
// }