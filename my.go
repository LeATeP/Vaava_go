package main

import (
	"fmt"
	"psql"
)

var result []map[string]any

func main() {
	p, _ := psql.PsqlConnect()
	selectItemsId := p.NewQuery("select * from items order by id;")
	updateItemsId := p.NewQuery("update items set amount = amount + $1 where id = 1;")

	result, _ = p.ExecQuery(selectItemsId)
	fmt.Println(result[0]["amount"])

	p.ExecCmd(updateItemsId, 100)
	
	result, _ = p.ExecQuery(selectItemsId)
	fmt.Println(result[0]["amount"])
}