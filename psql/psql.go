package psql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

// Psql_connect is main initialize fn, that connect to db and give interface
func PsqlConnect() (PsqlInterface, error) {
	var p Psql
	var err error
	config := init_config()
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.host, config.port, config.user, config.password, config.dbname)
	// prepare connecting to db
	p.Sql, err = sql.Open("postgres", psqlconn)
	if err != nil {
		return p, err
	}
	// checking connection if working
	err = p.Sql.Ping()
	if err != nil {
		return p, err
	}
	p.QueryMap = make(map[int64]MyQuery)
	var pA PsqlInterface = p
	return pA, nil
}

// func NewQuery accept sql cmd and return int identifier that used for Psql to find the right  Query
func (p Psql) NewQuery(s string) (int64, error) {
	var length int64 = int64(len(p.QueryMap) + len(p.ExecMap))
	prep, err := p.Sql.Prepare(s)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	p.QueryMap[length] = MyQuery{Query: s, PrepStmt: prep}
	return length, nil
}
func (p Psql) ExecQuery(i int64, args ...any) ([]map[string]any, error) {
	return p.QueryMap[i].runPrepQuery(args...)
}
func (p Psql) ExecCmd(i int64, args ...any) error {
	_, err := p.QueryMap[i].PrepStmt.Exec(args...)
	return err
}
func (p Psql) CloseQuery(i int64) {
	p.QueryMap[i].PrepStmt.Close()
}

func (q MyQuery) runPrepQuery(args ...any) ([]map[string]any, error) {
	maps := []map[string]any{}
	rowsStack := [][]any{}
	rows, err := q.PrepStmt.Query(args...)
	if err != nil {
		return nil, err
	}

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	length := len(columns)
	for rows.Next() {
		data, pointers := makePointers(length)
		if err := rows.Scan(pointers...); err != nil {
			log.Println(err)
			rows.Close()
		}
		rowsStack = append(rowsStack, data)
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
