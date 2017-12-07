package data

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type SqlDatabase struct {
	DB *sql.DB
}

func (db *SqlDatabase) Create(kind string, data ...interface{}) (string, error) {
	// TODO: INSERT

	return "", nil
}

func (db *SqlDatabase) Get(kind string, id string) (string, error) {
	// TODO: SELECT
	return "", nil
}

func (db *SqlDatabase) List(kind string, query ...interface{}) ([]string, error) {
	// TODO: SELECT
	return []string{}, nil
}

func (db *SqlDatabase) Update(kind string, id string, data ...interface{}) error {
	// TODO: UPDATE
	return nil
}

func (db *SqlDatabase) Delete(kind string, id string) error {
	// TODO: DELETE
	return nil
}

func NewSqlDatabase(protocol string, user string, password string, host string, port int, database string) SqlDatabase {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s(%s:%d)/%s", user, password, protocol, host, port, database))
	if err != nil {
		log.Fatalf("Couldn't connect to MySQL database: %v", err)
	}

	return SqlDatabase{DB: db,}
}