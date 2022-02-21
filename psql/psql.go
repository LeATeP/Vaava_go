package psql

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	ut "vaava/utils"

	_ "github.com/lib/pq"
)

var db *sql.DB
var (
	// hostname = os.Getenv("HOSTHAME")
    host     = os.Getenv("PSQL_HOST")
    dbname   = os.Getenv("PSQL_DB")
    user     = os.Getenv("PSQL_USER")
    password = os.Getenv("PGPASSWORD")
	port     = 5432
	err 	   error
)

func Psql_connect() error {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
							host, port, user, password, dbname)
	// Connecting to db
    db, err = sql.Open("postgres", psqlconn); 			if err != nil { return err }
	// checking connection if working
    err = db.Ping(); 									if err != nil { return err }

    fmt.Println("Connected!")
	return nil
}

func QuerySelect(sql_cmd string) ([]map[string]string, error) {
	rows, err 		:= db.Query(sql_cmd);				if err != nil {	return nil, err }
	defer rows.Close()

	columns, _ 		:= rows.Columns()
	rowsStack, err := iterRows(rows, len(columns));		if err != nil {	return nil, err }

	formedMap 		:= ut.ConvetIntoMap(rowsStack, columns)
	return formedMap, nil
}

func iterRows (rows *sql.Rows, lent int) (*[][]string, error) {
	rowsStack		:= [][]string{}

	for rows.Next() {
		content 	:= make([]string, lent)
		pointers 	:= ut.CreatePointers(&content)
	
		err := rows.Scan(pointers...);					if err != nil {	return nil, err }
		rowsStack 	= append(rowsStack, content)
	}
	err = rows.Err();									if err != nil {	return nil, err }
	return &rowsStack, nil
}

func Exec(sql_cmd string) error {
	_, err := db.Exec(sql_cmd);							if err != nil { return err }
	return nil
}
