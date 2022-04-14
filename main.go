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
	var err error
	res.Materials = map[string]int64{}
	client, err = server.NewClient()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Client Started")

	sendInfoToServer()
	go recvServer() // main loop, listening to server
	if !isServerReady() {
		log.Fatalln("Server is not ready")
		return
	}
	mining() // main production thread, generating loot
}
func isServerReady() bool { // 5 second to check if server have send information
	for i := 0; !client.FromServer.Running && i < 10; i++ {
		time.Sleep(time.Second / 2)
	}
	return client.FromServer.Running
}

func sendInfoToServer() {
	client.Send.Encode(&server.Message{MsgCode: 2, FromClient: client.FromClient})
}

func mining() {
	var err error
	defer preprareToShutDown()

	// starting main loop
	for ; client.FromClient.Running; time.Sleep(client.FromServer.TickSpeed) {
		err = client.Send.Encode(generateLoot())
		if err != nil {
			log.Printf("Can't sent msg: %v\n", err)
			return
		}
	}
}
func preprareToShutDown() {
	client.FromClient.Running = false
	client.Send.Encode(&server.Message{MsgCode: 4, FromClient: client.FromClient})
	client.Conn.Close()

}
func generateLoot() *server.Message {
	res.Materials = map[string]int64{}
	res.Materials["Rock"] += calculateChance(rockChance)
	fmt.Printf("+%v: Mined\n", res.Materials["Rock"])
	return &server.Message{MsgCode: 6, Resources: res}
}

func calculateChance(num int64) int64 {
	if rand.Int63n(num) == 0 {
		return 1
	}
	return 0
}

func recvServer() {
	var err error
	var msg server.Message
	defer preprareToShutDown()

	for client.FromClient.Running {
		msg = server.Message{}
		if err = client.Receive.Decode(&msg); err != nil {
			log.Printf("[Error in receiving msg]: %v", err)
			return
		}
		switch msg.MsgCode {
		case 1: // ping, saying that server is still alive
		case 2: // get info about the server
			client.FromServer = msg.FromServer
		case 3: // signal to change settings to...
		case 4: // signal to shutdown
			log.Fatalln("Signal to shutdown at")
			preprareToShutDown()
			return
		case 5: // signal to reload
		default:
			log.Println("0, something wrong")
		}
	}
}
