package main

import (
	// "fmt"
	// "math/rand"
	"fmt"
	"log"
	"math/rand"
	"time"
	"vaava/psql"
)
var server serverHandler

func main() {
	var err error
    // db, err = psql.Psql_connect()
    // if err != nil { log.Fatalln(err) }

	server, err 	= client()
    if err != nil { log.Fatalln(err) }

	for {
		num := rand.Int63n(10000)
		msg := fmt.Sprintf("client1: %v\n", num)
		err := server.sendMsg(msg)
		if err != nil {
			log.Printf("failed send msg: %v", err)
			break
		}
		time.Sleep(time.Second)
	}
	// query(db, "")
}

func query(db psql.DbInterface, cmd string) error {
    data, err := db.QuerySelect("select * from items;")
    if err != nil { log.Fatalln(err) }

    for _, w := range data {
		fmt.Printf("[%v] %-15v %d\n", w["id"], w["name"], w["amount"])
	}
	return nil
}