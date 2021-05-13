package data

import (
	"bytes"
	"io"
	"log"
)

func GetFile(id int, fileChan chan io.ReadSeeker) error {

	q := `SELECT file FROM uploadfile WHERE id=?1`
	db := getConnection()
	defer db.Close()
	row, err := db.Query(q, id)
	if err != nil {
		log.Println("error at line 8 in file uploadingFileGo/src/data/getFile.go")
		return err
	}
	var binaryFile []byte

	for row.Next() {
		if err := row.Scan(&binaryFile); err != nil {
			log.Println("error at line 21 in file uploadingFileGo/src/data/getFile.go")
			return err
		}

	}
	fileChan <- bytes.NewReader(binaryFile)

	return nil

}
