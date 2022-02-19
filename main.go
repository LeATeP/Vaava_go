package main

import (
	"fmt"
	ps "vaava/psql"
)

type asd struct {
	name int
}

func main() {
	ps.Psql_connect()
	x, _ := ps.QuerySelect("select id, name, amount from items")
	for _, r := range x {
		fmt.Println(r)
}
}
