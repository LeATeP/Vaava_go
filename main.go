package main

import (
	// "fmt"
	// "math/rand"
	"log"
	// "time"
	"vaava/psql"
)


func main() {
	db, err := psql.Psql_connect()
	if err != nil { log.Fatalln(err) }

	err = db.Exec("update items set amount = amount +1000 where name = 'Rock';")
	if err != nil { log.Fatalln(err) }
	
}
