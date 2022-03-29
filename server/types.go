package server

import (
	"encoding/gob"
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
	Status 		 string
	Start 		 time.Time
	TickDataSend time.Duration
	Running 	 bool
	Shutdown  	 bool
	ContainerId  string
}
type Client struct {
	Conn		 net.Conn
	Send		*gob.Encoder
	Receive		*gob.Decoder
	AboutClient *AboutClientInfo
}
type MsgFormat struct {
	MsgCode   int
	Resources Resources
	CInfo     AboutClientInfo
	SInfo 	  Info
}
type Resources struct {
	Materials map[string]uint64
}
// ---------------------------------------------
// Client Specific types