// package psql utils.go is for fn used in psql 
// but not necessary to be accessible outside of the package
package psql

import (
	"sort"
	"os"
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

func convetIntoMap(slices [][]string, columns []string) []map[string]string {
	newMaps 	:= []map[string]string{}

	for _, data := range slices {
		newMap 	:= map[string]string{}
		for i, colName := range columns {
			newMap[colName] = data[i]
		}
		newMaps = append(newMaps, newMap)
	}
	_, exist := newMaps[0]["id"]
	if exist {
  		sort.Slice(newMaps, func(i, j int) bool { return newMaps[i]["id"] < newMaps[j]["id"]})
	}
	return newMaps
}

func makePointers(rows_len int) ([]string, []any) {
	content  := make([]string, rows_len)
 	pointers := make([]any, rows_len)
	for i := range content {
		pointers[i] = &content[i]
	}
	return content, pointers

}
