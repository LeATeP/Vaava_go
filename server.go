package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"psql"
	"server"
	"time"
)

type connection struct {
	server *server.Server
	psql   psql.PsqlInterface
}

var (
	i int64              // client connection number
	res server.Resources // data client sending to process
	pool connection
)
const (
	selectTable = `select $1 from $2 order by id`
	updateItemAmount = `update items set amount = amount + $1 where id = $2;`
)
func main() {
	res.Materials = map[string]int64{}
	pool.server = server.NewServer()
	fmt.Println("server started")

	p, err := psql.PsqlConnect()
	if err != nil {
		log.Printf("%v\n", err)
	}
	pool.psql = p
	fmt.Println("psql conn ready")

	go pool.updateDB()
	for ; ; i++ {
		pool.AcceptConn(i)
	}

}
func (p *connection) updateDB() {
	var err error
	prep, err := p.psql.NewQuery(updateItemAmount)
	if err != nil {
		log.Printf("can't Run updateDB, %v\n", err)
		return
	}
	for ;; time.Sleep(time.Second) {
		for k, v := range res.Materials {
			err = p.psql.ExecCmd(prep, res.Materials[k], v)
			if err != nil {
				log.Printf("[Error in executing query]: %v", err)
				return
			}
			fmt.Println(k, v, "complete")
			res.Materials[k] = 0
		}
	}
}

func (c *connection) AcceptConn(i int64) {
	conn, err := c.server.Listener.Accept() // listen for clients
	if err != nil {
		log.Printf("[failed to connect]: %v\n", err)
	}
	fmt.Printf("connected [%v]: %v\n", i, conn)

	c.server.ClientConn[i] = &server.Client{
		Conn:    conn,
		Receive: gob.NewDecoder(conn),
		Send:    gob.NewEncoder(conn),
	}
	go c.ManageConnection(i)

}
func (c *connection) ManageConnection(i int64) {
	var err error
	var msg server.MsgFormat
	client := c.server.ClientConn[i]
	for {
		msg = server.MsgFormat{}
		if err = client.Receive.Decode(&msg); err != nil {
			log.Printf("%v [err]: %v\n", i, err) // well would be to put client identifiers like containerId and stuff
			client.Conn.Close()
			delete(c.server.ClientConn, i)
			return
		}
		switch msg.MsgCode {
		case 1: // get ping that client is active
		case 2: // get info about client
			client.AboutClient = msg.CInfo
		case 3: // something changed in client
		case 4: // client shutting down
			client.Conn.Close()
			delete(c.server.ClientConn, i)
			return
		case 5: // client reloading
		case 6: // update resources
			for i, k := range msg.Resources.Materials {
				res.Materials[i] += k
				fmt.Println(res.Materials)
			}
		default:
			fmt.Println("0, something wrong")
		}
	}
}

// c.Send.Encode
// c.Receive.Decode
// c.Conn.Close
// work just fine, but if needed to be put in interface, then it will be necessary?
