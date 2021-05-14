package data

import (
	"errors"
	"io/ioutil"
	"log"
	"mime/multipart"
)

// upload the file into the database
func UploadFile(file multipart.File, header *multipart.FileHeader) error {
	q := `INSERT INTO uploadfile( name,file ) VALUES(?1,?2)` // insert the name and  the file
	db := getConnection()                                    //get the connection
	defer db.Close()
	// convert the file multipart.File into a binary type
	fileBinary, _ := ioutil.ReadAll(file)
	stm, _ := db.Prepare(q)

	defer stm.Close()
	columns, err := stm.Exec(header.Filename, fileBinary)
	if err != nil {
		log.Println("err in file uploadingFileGo/src/data/uploadingFile.go line 21")
		return err
	}
	if r, _ := columns.RowsAffected(); r > 1 {
		return errors.New("how this could be happend")
	}
	return nil
}
