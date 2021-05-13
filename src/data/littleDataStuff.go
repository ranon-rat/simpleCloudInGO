package data

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func getConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "./database/database.db")
	if err != nil {
		log.Println("error in file uploadingFileGo/src/data/conn.go line 9\n", err)
	}
	return db
}
func GetSize() int {
	db := getConnection()
	m, _ := db.Query("SELECT COUNT(*) FROM uploadfile")
	many := 0
	for m.Next() {
		m.Scan(&many)
	}
	return many
}

// this check if the database have the same filename in the database
func Exist(header int) error {
	db := getConnection()
	r, _ := db.Query("SELECT COUNT(*) FROM uploadfile WHERE id=?1 ", header)
	howMany := 0
	for r.Next() {
		r.Scan(&howMany)
	}
	if howMany < 0 {
		return errors.New("your file already exist in the database")
	}
	return nil
}
