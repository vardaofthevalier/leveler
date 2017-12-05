package data

// import (
// 	"fmt"
// 	"log"
// 	"database/sql"
// 	_ "github.com/go-sql-driver/mysql"
// )

// type SqlDatabase struct {
// 	User string
// 	Password string
// 	Protocol string
// 	Host string
// 	Port int
// 	Database string
// }

// func (db *SqlDatabase) Create(kind string, keys map[string]interface, data string) (string, error) {
// 	// insert 
// }

// func (db *SqlDatabase) Get(kind string, id string) (string, error) {
// 	// select 
// }

// func (db *SqlDatabase) List(kind string, query string) ([]string, error) {
// 	// select with
// }

// func (db *SqlDatabase) Update(kind string, id string, data string) {
// 	// update
// }

// func (db *SqlDatabase) Delete(kind string, id string) error {
// 	// delete
// }

// func NewSqlDatabase(protocol string, user string, password string, host string, port int, database string) SqlDatabase {
// 	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s(%s:%d)/%s", user, password, protocol, host, port, database))
// 	if err != nil {
// 		log.Fatalf("Couldn't connect to MySQL database: %v", err)
// 	}

// 	return db
// }