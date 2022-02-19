package psql

import (
    "database/sql"
    "fmt"
	"os"
    _ "github.com/lib/pq"
	ut "vaava/utils"
)

var db *sql.DB
var (
	// hostname = os.Getenv("HOSTHAME")
    host     = os.Getenv("PSQL_HOST")
    dbname   = os.Getenv("PSQL_DB")
    user     = os.Getenv("PSQL_USER")
    password = os.Getenv("PGPASSWORD")
	port     = 5432
)

func Psql_connect() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
							host, port, user, password, dbname)
	var err error
    db, err = sql.Open("postgres", psqlconn) // connecting to db

	ut.CheckError(err, "failed open connection sql.Open")

    err = db.Ping() // check connection

    ut.CheckError(err, "failed ping")
    fmt.Println("Connected!")
}

func QuerySelect(sql_cmd string) ([]map[string]string, error) {
	rows, err 		:= db.Query(sql_cmd)
	ut.CheckError(err, "attempt to query db.Query")
	defer rows.Close()

	columns, _ 		:= rows.Columns()
	rowsStack := iterRows(rows, len(columns))

	formedMap 		:= ut.ConvetIntoMap(rowsStack, columns)
	return formedMap, nil
}

func iterRows (rows *sql.Rows, lent int) *[][]string {
	var err error
	rowsStack		:= [][]string{}

	for rows.Next() {
		content 	:= make([]string, lent)
		pointers 	:= ut.CreatePointers(&content)
	
		err := rows.Scan(pointers...)
		ut.CheckError(err, "attempt to Scan rows.Next")
		
		rowsStack 	= append(rowsStack, content)
	}
	err = rows.Err()
	ut.CheckError(err, "attempt ending rows.Err")

	return &rowsStack
}

func Exec(sql_cmd string) bool {
	_, err := db.Exec(sql_cmd)
	return err == nil
}
