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
	dropChance map[string]int64 = map[string]int64{
		"Rock": 5,
	}
	maxDropAmount map[string]int64 = map[string]int64{
		"Rock": 15,
	}
)

func main() {
	rand.Seed(time.Now().UnixMicro())
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

func mining() {
	defer preprareToShutDown()
	var err error

	// starting main loop
	for ; client.FromClient.Running ; time.Sleep(client.FromServer.TickSpeed) {
		err = client.Send.Encode(generateLoot())
		if err != nil {
			log.Printf("Can't sent msg: %v\n", err)
			return
		}
	}
}
func generateLoot() *server.Message {
	dropList := []string{"Rock"}
	res.Materials = map[string]int64{}
	for _, k := range dropList {
		res.Materials[k] += multipleChances(k)
	}
	fmt.Printf("+%v: Mined\n", res.Materials["Rock"])
	return &server.Message{MsgCode: 6, Resources: res}
}

func multipleChances(name string) (total int64) {
	for i := int64(0); i < maxDropAmount[name]; i++ {
		total += drop(dropChance[name])
	}
	return
}

func drop(chance int64) int64 {
	if chance < 2 {
		return 1
	}
	if rand.Int63n(chance) == 0 {
		return 1
	}
	return 0
}
