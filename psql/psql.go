package psql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

var (
	psql  *sql.DB
	mutex = sync.RWMutex{}
)

// Psql_connect is main initialize fn, that connect to db and give interface
func PsqlConnect() (PsqlInterface, error) {
	var cmd CmdMap
	var err error
	config := init_config()
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.host, config.port, config.user, config.password, config.dbname)
	// prepare connecting to db
	psql, err = sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}
	// checking connection if working
	err = psql.Ping()
	if err != nil {
		return nil, err
	}
	cmd.QueryMap = make(map[int64]MyQuery)
	var pA PsqlInterface = cmd
	return pA, nil
}

// func NewQuery accept sql cmd and return int identifier that used for Psql to find the right  Query
func (p CmdMap) NewQuery(cmd string) (id int64, err error) {
	mutex.Lock()
	id = int64(len(p.QueryMap) + len(p.ExecMap))
	mutex.Unlock()

	prep, err := psql.Prepare(cmd)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	mutex.Lock()
	p.QueryMap[id] = MyQuery{Query: cmd, PrepStmt: prep}
	mutex.Unlock()
	return id, nil
}
func (p CmdMap) ExecQuery(i int64, args ...any) ([]map[string]any, error) {
	result, err := p.QueryMap[i].runPrepQuery(args...)
	return result, err
}
func (p CmdMap) ExecCmd(i int64, args ...any) error {
	mutex.Lock()
	_, err := p.QueryMap[i].PrepStmt.Exec(args...)
	mutex.Unlock()
	return err
}
func (p CmdMap) CloseQuery(i int64) {
	mutex.Lock()
	p.QueryMap[i].PrepStmt.Close()
	delete(p.QueryMap, i)
	mutex.Unlock()
}

// fast query that prepare and closing after completing exec query, returning only error
func (p CmdMap) ExecCmdFast(cmd string, args ...any) {
	id, _ := p.NewQuery(cmd)
	p.ExecCmd(id, args...)
	p.CloseQuery(id)
}

func (q MyQuery) runPrepQuery(args ...any) ([]map[string]any, error) {
	rowsStack := [][]any{}
	mutex.Lock()
	rows, err := q.PrepStmt.Query(args...)
	mutex.Unlock()
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
		return nil, fmt.Errorf("0 rows was received")
	}
	maps := convetIntoMap(rowsStack, columns)
	return maps, nil
}
