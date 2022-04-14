package unitcreate

import (
	"fmt"
	"time"
	"log"
	"vaava/psql"
)

func main() {
	psql.Psql_connect()

	cmd := ""
	var err error
	xp, total := 0, 0
	for i:=2; i < 1000; i++ {
		time.Sleep(time.Second / 15)
		xp = i * i
		total += xp
		
		cmd = fmt.Sprintf("insert into levels(level, xp, total) values (%v, %v, %v);", i, xp, total)
		err = psql.Exec(cmd)
		if err != nil {	log.Fatal(err) }

		fmt.Println(i, xp, total)
	}
}