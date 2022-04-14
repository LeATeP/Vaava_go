package main

import (
	"fmt"
	"psql"
)


var result []map[string]any
var a int64 = 10
func main() {
	p, _ := psql.PsqlConnect()
	// selectItemsId, _ := p.NewQuery("select * from items order by id;")
	// updateItemsId := p.NewQuery("update items set amount = amount + $1 where id = 1;")
	// data, _ := p.ExecQuery(selectItemsId, )
	// fmt.Println(data[0])
}
