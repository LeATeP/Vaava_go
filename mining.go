package main

import (
	"fmt"
	"os"
	"time"
	drop "vaava/package/drop"
	ps "vaava/psql"
)
// the first mining script will do is to pull drop chances from sql

func main() {
	ps.Psql_connect()

	fmt.Println("starting mining")
	tm := 1 * time.Second
	for {
		time.Sleep(tm / 10)
		// ps.Exec("update items set amount = amount + 1 where id = 1;")
		drop := drop.GenerateDrop()

		for name, amount := range drop {
			cmd := fmt.Sprintf("update items set amount = amount + %v where name = '%v';", amount, name)
			ps.Exec(cmd)
		}
	}
}

func setContainerId() {
	hostname := os.Getenv("HOSTHAME")
	x := fmt.Sprintf("%v", hostname)
	ps.Exec(x)
}