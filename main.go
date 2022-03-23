package main

import (
	// "fmt"
	// "math/rand"
	"fmt"
	"log"
	// "time"
	"vaava/psql"
)


func main() {
    db, err := psql.Psql_connect()
    if err != nil { log.Fatalln(err) }

    data, err := db.QuerySelect("select * from items;")
    if err != nil { log.Fatalln(err) }

    for _, w := range data {
        fmt.Printf("[%v] %v: %v\n", w["id"], w["name"], w["amount"])
    }

}
