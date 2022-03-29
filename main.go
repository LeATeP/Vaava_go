package main

import (
	"fmt"
	"log"
	"time"
	"vaava/psql"
	"vaava/server"
)

func main() {
	srv := server.NewClient()
	
	time.Sleep(time.Second)
	srv.Send.Encode(&server.MsgFormat{MsgCode: "01", Name: "asd", Num: 100})
	
}



















func query(db psql.DbInterface, cmd string) error {
    data, err := db.QuerySelect("select * from items;")
    if err != nil { log.Fatalln(err) }

    for _, w := range data {
		fmt.Printf("[%v] %-15v %d\n", w["id"], w["name"], w["amount"])
	}
	return nil
}