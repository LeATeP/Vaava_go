package server

import (
	"encoding/gob"
	"math/rand"
	"net"
	"time"
)
const (
	networkTick 	   = time.Second
	listenToLocal      = "localhost:9000"
	connectToLocal     = "localhost:9000"
	connectToNetwork   = "postgres:9000"
)
// Server Specific Types "Below"
// type Info about the server that passes to client
type Info struct {
	Id			 string
	Name 		 string
	ContainerId  string
	MaxLoad 	 int
}
type Server struct {
	Start 		 time.Time
	Info 		 Info
	Stats 		*Stats
	Running 	 bool
	Shutdown  	 bool
	Reloading	 bool
	Listener 	 net.Listener
	ClientConn   map[uint64]*Client  	  // info the server have about client
}
type Stats struct {
	InMsgs	  	 int64
	OutMsgs      int64
	InBytes      int64
	OutBytes     int64
}
type AboutClientInfo struct {
	Id 			 int
	Name 		 string
	Doing 		 string
	Start 		 time.Time
	Running 	 bool
	Shutdown  	 bool
	ContainerId  string
	RandSeed 	*rand.Rand
}
type Client struct {
	Conn		 net.Conn
	Send		*gob.Encoder
	Receive		*gob.Decoder
	AboutClient *AboutClientInfo
}
type MsgFormat struct {
	MsgCode string
	Name 	string
	Num 	int
}
type Resources struct {
	Coins int
}
// ---------------------------------------------
// Client Specific types