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
var res server.Resources
var (
	rockChance int64 = 1
)

func main() {
	res.Materials = map[string]int64{}
	fmt.Println("Client Started")
	time.Sleep(time.Second / 2) // testing

	client = server.NewClient()
	sendInfoToServer()
	go mining()  // main production thread, generating loot
	recvServer() // main loop, listening to server
}
func sendInfoToServer() {
	client.Send.Encode(&server.MsgFormat{MsgCode: 2, CInfo: client.AboutClient})
}

func mining() {
	var err error
	var sleep time.Duration = time.Second
	defer preprareToShutDown()

	// starting main loop
	for ; client.AboutClient.Running; time.Sleep(sleep) {
		err = client.Send.Encode(generateLoot())
		if err != nil {
			log.Printf("Can't sent msg: %v\n", err)
			return
		}
	}
}
func preprareToShutDown() {
	client.Send.Encode(&server.MsgFormat{MsgCode: 4})
	client.Conn.Close()
	client.AboutClient.Running = false
}
func generateLoot() *server.MsgFormat {
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

func recvServer() {
	var err error
	var msg server.MsgFormat
	defer preprareToShutDown()

	for client.AboutClient.Running {
		msg = server.MsgFormat{}
		if err = client.Receive.Decode(&msg); err != nil {
			log.Printf("[Error in receiving msg]: %v", err)
			return
		}
		switch msg.MsgCode {
		case 1: // ping, saying that server is still alive
		case 2: // get info about the server
		case 3: // signal to change settings to...
		case 4: // signal to shutdown
			log.Fatalln("Signal to shutdown at")
			return
		case 5: // signal to reload
		default:
			log.Println("0, something wrong")
		}
	}
}
