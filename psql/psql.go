// package psql is mainly to give interface to manage psql db query's and stuff
package psql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Psql_connect is main initialize fn, that connect to db and give interface
func Psql_connect() (d DBStruct, err error) {
	config 		:= init_config()
	psqlconn 	:= fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
							config.host, config.port, config.user, config.password, config.dbname)
	
	// prepare connecting to db
    d.DB, err = sql.Open("postgres", psqlconn)
	if err != nil {
		return d, err
	}
	// checking connection if working
    if err = d.DB.Ping(); err != nil {
		return d, err
	}
	return d, err
}
func (d *DBStruct) InnateQuery(query string, tableName string) ([]TableItems, error) {
	rows, err := d.DB.Query(query)
	if err != nil {
		log.Printf("[Err in Query]: %v\n", err)
		return []TableItems{}, nil
	}
	sliceData, err := iterRows(rows, tableName)
	if err != nil {
		log.Printf("[Err in iterRows] --- %v\n", err)
	}
	return sliceData, nil

}
func iterRows(rows *sql.Rows, tableName string) ([]TableItems, error) {
	sliceData := []TableItems{}
	columns, _   := rows.Columns()
	defer rows.Close()
	for rows.Next() {
		data, pointers := itemsPointers(columns, tableName)
		if err := rows.Scan(pointers...); err != nil {
			log.Printf("[Err in iter of rows]: %v\n", err)
			rows.Close()
		}
		sliceData = append(sliceData, *data)
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		log.Printf("[rows.Err]: %v\n", err)
		return []TableItems{}, err
	}
	return sliceData, nil
}
func itemsPointers(col []string, tableName string) (*TableItems, []any) {
	data := new(TableItems)
	p    := []any{}
	item := map[string]any{
		"id": 			&data.Id,
		"name": 		&data.Name,
		"amount": 		&data.Amount,
		"amount_limit": &data.MaxAmount,
	}
	for _, colName := range col {
		value, exist := item[colName]
		if exist {
			p = append(p, value)
		}
	}
	return data, p
}