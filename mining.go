package main

import (
	"fmt"
	"os"
	"time"
	"log"
	drop "vaava/package/drop"
	ps "vaava/psql"
)
// the first mining script will do is to pull drop chances from sql
var (
	err error
)


func main() {
	err = ps.Psql_connect()
	if err != nil {	log.Fatal(err) }

	fmt.Println("starting mining")
	tm := time.Second / 10
	for {
		time.Sleep(tm)
		// ps.Exec("update items set amount = amount + 1 where id = 1;")
		drop := drop.GenerateDrop()
		fmt.Println(drop)

		for name, amount := range drop {
			cmd := fmt.Sprintf("update items set amount = amount + %v where name = '%v';", amount, name)
			err = ps.Exec(cmd)
			if err != nil {	log.Fatal(err) }
		}
	}
}

func setContainerId() {
	hostname := os.Getenv("HOSTHAME")
	x := fmt.Sprintf("%v", hostname)
	ps.Exec(x)
}