// package psql utils.go is for fn used in psql
// but not necessary to be accessible outside of the package
package psql

import (
	"database/sql"
	"os"
	"sort"
)

// con_config is connection config
type con_config struct {
	hostname string
    host     string
    dbname   string
    user     string
    password string
	port     int
}
type TableItems struct {
	Id  	  int
	Name 	  string
	Amount    int
	MaxAmount int
}
// type QueryStruct is to hold prep Query's and description of query's 
type PrepQuerySelect struct {
	QueryName 	string   // Key for the map DBStruct
	Description string   // what Query is doing
	TypeOfQuery string   // like select/update/delete/insert/
	PrepQuery  *sql.Stmt // the prep itself
	Query 		string   // the cmd itself that used for prep
	TableName 	string   
}
// DBStruct is to be used as a `connection` handler, for method to use
type Psql struct {
	Sql *sql.DB
	QueryMap map[string]PrepQuerySelect  // key: Table Name 

}
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

func convetIntoMap(slices [][]any, columns []string) []map[string]any {
	newMaps 	:= make([]map[string]any, len(slices))

	for i, data := range slices {
		newMap 	:= map[string]any{}
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
  		sort.Slice(newMaps, func(i, j int) bool { return newMaps[i]["id"].(int64) < newMaps[j]["id"].(int64)})
	}
}

func makePointers(rows_len int) ([]any, []any) {
	content  := make([]any, rows_len)
 	pointers := make([]any, rows_len)
	for i := range content {
		pointers[i] = &content[i]
	}
	return content, pointers

}