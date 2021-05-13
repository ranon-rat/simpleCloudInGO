package data

import (
	"bytes"
	"io"
	"log"
)

func GetFile(id int, fileChan chan io.ReadSeeker, nameChan chan string) error {

	q := `SELECT (file,name) FROM uploadfile WHERE id=?1`
	db := getConnection()
	defer db.Close()
	row, err := db.Query(q, id)
	if err != nil {
		log.Println("error at line 8 in file uploadingFileGo/src/data/getFile.go")
		return err
	}
	var binaryFile []byte
	var nameStr string
	for row.Next() {
		if err := row.Scan(&binaryFile, &nameStr); err != nil {
			log.Println("error at line 21 in file uploadingFileGo/src/data/getFile.go")
			return err
		}

	}
	fileChan <- bytes.NewReader(binaryFile)
	nameChan <- nameStr
	return nil

}
