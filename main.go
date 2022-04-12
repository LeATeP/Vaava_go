package main

import (
	"fmt"
	"log"
	"math/rand"
	"server"
	"time"
)

// implementation of mining
// 1. connect to a server
// 2. [wait mode], wait for signal 2, info about unit, or what to do
// 3. start independent loop to listen to server, for change unit info or shutdown
// 4. main loop is unit operation, sending info req by server to operate and manage unit
var client *server.Client
var rockChance int64 = 1
var res server.Resources

func main() {
	res.Materials = map[string]int64{}
	fmt.Println("Client Started")
	time.Sleep(time.Second / 2) // testing

	client = server.NewClient()
	client.Send.Encode(&server.MsgFormat{MsgCode: 2, CInfo: client.AboutClient})
	go Mining()
	RecvServer()
}

func Mining() {
	var err error
	sleep := time.Second

	// starting main loop
	for ; client.AboutClient.Running; time.Sleep(sleep) {
		err = client.Send.Encode(genLoot())
		if err != nil {
			log.Printf("Can't sent msg: %v\n", err)
			client.Conn.Close()
			client.AboutClient.Running = false
			return
		}
	}
}
func genLoot() *server.MsgFormat {
	res.Materials = map[string]int64{}
	res.Materials["Rock"] += calculateChance(rockChance)
	fmt.Printf("+%v: Mined\n", res.Materials["Rock"])
	return &server.MsgFormat{MsgCode: 6, Resources: res}
}

func calculateChance(num int64) int64 {
	if rand.Int63n(num) == 0 {
		return 1
	}
	return 0
}

func RecvServer() {
	var err error
	var msg server.MsgFormat
	for client.AboutClient.Running {
		msg = server.MsgFormat{}
		if err = client.Receive.Decode(&msg); err != nil {
			log.Printf("[Error in receiving msg]: %v", err)
			client.Conn.Close()
			client.AboutClient.Running = false
			return
		}
		switch msg.MsgCode {
		case 1: // ping, saying that server is still alive
		case 2: // get info about the server

		case 3: // signal to change settings to...
		case 4: // signal to shutdown
			log.Fatalln("Signal to shutdown at")
			client.Conn.Close()
			client.AboutClient.Running = false
			return
		case 5: // signal to reload
		default:
			log.Println("0, something wrong")
		}
	}
}
