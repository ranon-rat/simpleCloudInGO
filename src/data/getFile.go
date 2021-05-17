package data

import (
	"bytes"
	"io"
)

func GetFile(id int, fileChan chan io.ReadSeeker) error {

	q := `SELECT file FROM uploadfile WHERE id=?1`
	db := getConnection()
	defer db.Close()

	binaryFile := new([]byte) //= []byte{14, 14, 14}
	db.QueryRow(q, id).Scan(binaryFile)

	fileChan <- bytes.NewReader(*binaryFile)

	return nil

}
