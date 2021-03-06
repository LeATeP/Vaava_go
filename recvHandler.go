package main

import (
	"fmt"
	"log"
	"server"
	"time"
)

func recvServer() {
	defer preprareToShutDown("recvServer")
	var err error
	var msg server.Message

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
			client.FromClient.Running = msg.FromServer.Running
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
func preprareToShutDown(name string) {
	fmt.Printf("Shutting down at %v", name)
	client.FromClient.Running = false
	client.Send.Encode(&server.Message{MsgCode: 4, FromClient: client.FromClient})
	client.Conn.Close()
}
func isServerReady() bool { // 5 second to check if server have send information
	for i := 0; !client.FromServer.Running && i < 10; i++ {
		time.Sleep(time.Second / 2)
	}
	fmt.Println(client.FromServer)
	return client.FromServer.Running
}
func sendInfoToServer() {
	client.Send.Encode(&server.Message{MsgCode: 2, FromClient: client.FromClient})
}

