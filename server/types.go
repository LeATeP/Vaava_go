package server

import (
	"encoding/gob"
	"net"
	"time"
)

const (
	networkTick      = time.Second
	listenToLocal    = "localhost:9000"
	connectToLocal   = "localhost:9000"
	connectToNetwork = "postgres:9000"
)

// Server Specific Types "Below"
// type Info about the server that passes to client
type Info struct {
	Id          string
	Name        string
	ContainerId string
	MaxLoad     int64
}
type Server struct {
	Start      time.Time
	Info       Info
	Stats      *Stats
	Running    bool
	Shutdown   bool
	Reloading  bool
	Listener   net.Listener
	ClientConn map[int64]*Client // info the server have about client
}
type Stats struct {
	InMsgs   int64
	OutMsgs  int64
	InBytes  int64
	OutBytes int64
}
type AboutClientInfo struct {
	Id           int64
	Start        time.Time
	TickDataSend time.Duration
	Running      bool
	Shutdown     bool
	ContainerId  string
	Unit         UnitInfo
}
type Client struct {
	Conn        net.Conn
	Send        *gob.Encoder
	Receive     *gob.Decoder
	AboutClient *AboutClientInfo
}
type UnitInfo struct {
	Id       int64
	Name     string
	Status   string
	Health   int64
	Strength int64
}
type MsgFormat struct {
	MsgCode   int64
	Resources Resources
	CInfo     *AboutClientInfo
	SInfo     Info
}
type Resources struct {
	Materials map[string]int64
}

// ---------------------------------------------
// Client Specific types
