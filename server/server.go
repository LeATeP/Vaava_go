package server

import (
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