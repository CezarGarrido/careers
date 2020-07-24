package driver

import (
	"database/sql"

	_ "github.com/lib/pq"
)

//DB: expose all database suport
type DB struct {
	SQL   *sql.DB
	// Mgo *mgo.database
}

var dbConn = &DB{}

func OpenPostgres(url string) (*DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	dbConn.SQL = db
	return dbConn, err
}
