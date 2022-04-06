// package psql is mainly to give interface to manage psql db query's and stuff
package psql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Psql_connect is main initialize fn, that connect to db and give interface
func Psql_connect() (p Psql, err error) {
	config 	  := init_config()
	psqlconn  := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
							config.host, config.port, config.user, config.password, config.dbname)
	// prepare connecting to db
    p.Sql, err = sql.Open("postgres", psqlconn)
	if err != nil { 
		return p, err
	}
	// checking connection if working
    err 	= p.Sql.Ping()
	if err != nil {
		return p, err
	}
	return p, nil
}
// func MakeQuery is good
// func (sql *Psql) MakeQuery(q PrepQuerySelect) {}
func RunStmtQuery(q *sql.Stmt) ([]map[string]any, error) {
	maps 		 := []map[string]any{}
	rowsStack	 := [][]any{}
	rows, err    := q.Query()
	if err  	 != nil {
		return nil, err
	}

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	length 		 := len(columns)
	for rows.Next() {
		data, pointers := makePointers(length)
		if err 		   := rows.Scan(pointers...); err != nil {
			log.Println(err)
			rows.Close()
		}
		rowsStack 	= append(rowsStack, data)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	if len(rowsStack) == 0 {
		return maps, fmt.Errorf("0 rows was received")
	}
	maps = convetIntoMap(rowsStack, columns)
	return maps, nil
}
func handleRows(r *sql.Rows) {}















// func (d *DBStruct) InnateQuery(query string, tableName string) ([]ItemsTable, error) {
	// rows, err := d.DB.Query(query)
	// if err != nil {
		// log.Printf("[Err in Query]: %v\n", err)
		// return []ItemsTable{}, nil
	// }
	// sliceData, err := iterRows(rows, tableName)
	// if err != nil {
		// log.Printf("[Err in iterRows] --- %v\n", err)
	// }
	// return sliceData, nil
// 
// }
// func iterRows(rows *sql.Rows, tableName string) ([]ItemsTable, error) {
	// sliceData := []ItemsTable{}
	// columns, _   := rows.Columns()
	// defer rows.Close()
	// for rows.Next() {
		// data, pointers := itemsPointers(columns, tableName)
		// if err := rows.Scan(pointers...); err != nil {
			// log.Printf("[Err in iter of rows]: %v\n", err)
			// rows.Close()
		// }
		// sliceData = append(sliceData, *data)
	// }
	// rows.Close()
	// if err := rows.Err(); err != nil {
		// log.Printf("[rows.Err]: %v\n", err)
		// return []ItemsTable{}, err
	// }
	// return sliceData, nil
// }
// func itemsPointers(col []string, tableName string) (*ItemsTable, []any) {
	// data := new(Items)
	// p    := []any{}
	// item := map[string]any{
		// "id": 			&data.Id,
		// "name": 		&data.Name,
		// "amount": 		&data.Amount,
		// "amount_limit": &data.MaxAmount,
	// }
	// for _, colName := range col {
		// value, exist := item[colName]
		// if exist {
			// p = append(p, value)
		// }
	// }
	// return data, p
// }