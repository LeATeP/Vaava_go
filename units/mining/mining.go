package mining

import (
	"fmt"
	"os"
	"time"
	"log"
	drop "vaava/drop"
	"vaava/psql"
)
// the first mining script will do is to pull drop chances from sql
var (
	err error
)

func main() {
	err = psql.Psql_connect()
	if err != nil {	log.Fatal(err) }

	fmt.Println("starting mining")
	tm := time.Second
	count, xp := 0, 0
	for {
		count ++
		time.Sleep(tm)
		// ps.Exec("update items set amount = amount + 1 where id = 1;")
		drop := drop.GenerateDrop()
		fmt.Println(drop)

		for name, amount := range drop {
			cmd := fmt.Sprintf("update items set amount = amount + %v where name = '%v';", amount, name)
			err = psql.Exec(cmd)
			if err != nil {	log.Fatal(err) }
			xp += amount
		}
		if count == 10 {
			count = 0
			cmd := fmt.Sprintf("update unit set xp = xp + %v where name = '%v'", xp, "Miner")
			err = psql.Exec(cmd)
			if err != nil {	log.Fatal(err) }
			
			xp = 0
		}
	}
}

func setContainerId() {
	hostname := os.Getenv("HOSTHAME")
	x := fmt.Sprintf("%v", hostname)
	psql.Exec(x)
}
