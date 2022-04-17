package psql

import (
	"database/sql"
	"os"
	"sort"
)

// Type con_config is connection config to database
type con_config struct {
	hostname string // name of the machine / service / containerID
	host     string // IP address of db or name of the network
	dbname   string // name of db to connect to
	user     string // db user
	password string // db password
	port     int    // port of db
}

// type QueryStruct is to hold prep Query's and description of query's
type MyQuery struct {
	// QueryName 	string   // Key for the map DBStruct
	// Description string   // what Query is doing
	// TypeOfQuery string   // like select/update/delete/insert/  (can just take [0] of the string of the query)
	PrepStmt *sql.Stmt // the prep itself
	Query    string    // the cmd itself that used for prep
	// TableName 	string
}


// DBStruct is to be used as a `connection` handler, for method to use
type CmdMap struct {
	QueryMap map[int64]MyQuery // key: Table Name
	ExecMap  map[int64]MyQuery // key: Table Name
}
type PsqlInterface interface {
	NewQuery(s string) (int64, error)
	ExecQuery(i int64, args ...any) ([]map[string]any, error)
	ExecCmd(i int64, args ...any) error
	CloseQuery(i int64)
	ExecCmdFast(cmd string, args ...any)  
}

func init_config() *con_config {
	return &con_config{
		hostname: os.Getenv("HOSTHAME"),
		host:     os.Getenv("PSQL_HOST"),
		dbname:   os.Getenv("PSQL_DB"),
		user:     os.Getenv("PSQL_USER"),
		password: os.Getenv("PGPASSWORD"),
		port:     5432,
	}
}

func convetIntoMap(slices [][]any, columns []string) []map[string]any {
	newMaps := make([]map[string]any, len(slices))

	for i, data := range slices {
		newMap := map[string]any{}
		for r, colName := range columns {
			newMap[colName] = data[r]
		}
		newMaps[i] = newMap
	}
	sortSliceOfMap(newMaps)
	return newMaps
}

func sortSliceOfMap(newMaps []map[string]any) {
	if len(newMaps) == 0 {
		return
	}
	_, exist := newMaps[0]["id"].(int64)
	if exist {
		sort.Slice(newMaps, func(i, j int) bool { return newMaps[i]["id"].(int64) < newMaps[j]["id"].(int64) })
	}
}

func makePointers(rows_len int) ([]any, []any) {
	content := make([]any, rows_len)
	pointers := make([]any, rows_len)
	for i := range content {
		pointers[i] = &content[i]
	}
	return content, pointers

}
