package psql

import (
	"database/sql"
	"fmt"
	"os"
	ut "vaava/utils"
	_ "github.com/lib/pq"
)

type dbStruct struct {
	conn *sql.DB
}

type dbInterface interface {
	Exec(string) error
	Ping() error
	QuerySelect(string) ([]map[string]string, error)
}

type con_config struct {
	hostname string
    host     string
    dbname   string
    user     string
    password string
	port     int
}
var (
	d dbStruct
	err error
)

func init_config() *con_config {
	return &con_config{
		hostname:   os.Getenv("HOSTHAME"),
		host: 		os.Getenv("PSQL_HOST"),
  		dbname: 	os.Getenv("PSQL_DB"),
  		user: 		os.Getenv("PSQL_USER"),
  		password:   os.Getenv("PGPASSWORD"),
		port: 		5432,
	}
}

func Psql_connect() (dbInterface, error) {
	config 		:= init_config()
	psqlconn 	:= fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
							config.host, config.port, config.user, config.password, config.dbname)


	
	// Connecting to db
    d.conn, err = sql.Open("postgres", psqlconn)			
	if err	   	!= nil { return nil, err }


	// checking connection if working
    err 		= d.Ping(); 									
	if err 		!= nil { return nil, err }
    fmt.Println("Connected!")

	content, pointers := MakePointers(5)
	fmt.Println(content, pointers)
	return dbInterface(&d), nil
}

func (d *dbStruct) QuerySelect(sql_cmd string) ([]map[string]string, error) {
	rows, err 		:= d.conn.Query(sql_cmd)
	if err != nil {	return nil, err }
	defer rows.Close()

	columns, err 	:= rows.Columns()
	if err != nil {	return nil, err }

	rowsStack, err  := iterRows(rows, len(columns))
	if err != nil {	return nil, err }

	formedMap 		:= ut.ConvetIntoMap(rowsStack, columns)
	return formedMap, nil
}

func iterRows (rows *sql.Rows, row_length int) ([][]string, error) {
	rowsStack		:= [][]string{}

	for rows.Next() {
		content, pointers := MakePointers(row_length)

		err := rows.Scan(pointers...)
		if err != nil {	return nil, err }
		
		rowsStack 	= append(rowsStack, content)
	}
	err = rows.Err()
	if err != nil {	return nil, err }
	
	return rowsStack, nil
}
func MakePointers(rows_len int) ([]string, []interface{}) {
	content  := make([]string, rows_len)
 	pointers := make([]interface{}, rows_len)
	for i := range content {
		pointers[i] = &content[i]
	}
	return content, pointers

}

func (d *dbStruct) Exec(sql_cmd string) error {
	_, err := d.conn.Exec(sql_cmd)					
	if err != nil { return err }
	return nil
}

func (d *dbStruct) Ping() error {
	return d.conn.Ping()
}