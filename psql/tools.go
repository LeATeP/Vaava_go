// package psql utils.go is for fn used in psql
// but not necessary to be accessible outside of the package
package psql

import (
	"database/sql"
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
type TableItems struct {
	Id  	  int
	Name 	  string
	Amount    int
	MaxAmount int
}
// type QueryStruct is to hold prep Query's and description of query's 
type PrepQuerySelect struct {
	Description string   // what Query is doing
	TypeOfQuery string   // like select/update/delete/insert/
	PrepQuery  *sql.Stmt // the prep itself
	Query 		string   // the cmd itself that used for prep
	TableName 	string   
}
// DBStruct is to be used as a `connection` handler, for method to use
type DBStruct struct {
	DB *sql.DB
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