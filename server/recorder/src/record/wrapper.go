package record

import (
	// "bufio"
	"database/sql"
	"fmt"
	"os"
	// "strings"
	//"sync"

	_ "github.com/lib/pq"
)

type dbconn interface {
	Insert(string)
	Select() string
}

type pgdb struct {
	db *sql.DB
}

func (pg pgdb) Insert(reqJson string) {
	pg.db.Exec(fmt.Sprintf(
`INSERT INTO reqData
	(reqJson)
	VALUES ('%s')`,
reqJson))
}

func (pg pgdb) Select() string {
	var reqJson string
	iter, err := pg.db.Query(
"SELECT reqJson FROM reqData")
	defer iter.Close()
	check(err)
	for iter.Next() {
		if err := iter.Scan(&reqJson); err != nil {
			check(err)
		}
	}
	return reqJson
}

func SetupDB(conn dbconn) {
	pg, _ := conn.(pgdb)
	_, err := pg.db.Exec(
`CREATE TABLE IF NOT EXISTS reqData (reqJson JSON)`)
	check(err)
}

func NewDBConn() dbconn {
	return NewPGDB()
}

func NewPGDB() pgdb {
	session, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		os.Getenv("USER"),
		os.Getenv("PASSWORD"),
		os.Getenv("DBNAME"),
		os.Getenv("HOST"),
		os.Getenv("DB_PORT"),
	))
	check(err)
	return pgdb{db: session}
}
