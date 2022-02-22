package main

import (
	"fmt"
	"os"
	"time"
	"log"
	drop "vaava/package/drop"
	"vaava/psql"
)
// the first mining script will do is to pull drop chances from sql
var (
	err error
)

func asd(y []int) {
	y[0] = 10
}

func main() {
	x := []int{1,2,3}
	asd(x)
	fmt.Println(x)


	err = psql.Psql_connect()
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
			err = psql.Exec(cmd)
			if err != nil {	log.Fatal(err) }
		}
	}
}

func setContainerId() {
	hostname := os.Getenv("HOSTHAME")
	x := fmt.Sprintf("%v", hostname)
	psql.Exec(x)
}
