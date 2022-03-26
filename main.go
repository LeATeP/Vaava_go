package main

import (
	// "fmt"
	// "math/rand"
	"fmt"
	"log"

	// "time"
	"vaava/psql"
)

var db psql.DbInterface

func main() {
	var err error
    db, err = psql.Psql_connect()
    if err != nil { log.Fatalln(err) }

	query("")

}

func query(cmd string) error {
    data, err := db.QuerySelect("select * from items;")
    if err != nil { log.Fatalln(err) }

    for _, w := range data {
		fmt.Printf("[%v] %-15v %d\n", w["id"], w["name"], w["amount"])
	}
	return nil
}