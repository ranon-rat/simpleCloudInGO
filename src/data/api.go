package data

import (
	"log"

	"github.com/ranon-rat/simpleCloudInGO/src/stuff"
)

func GetSize() int {
	db := getConnection()
	defer db.Close()

	m, _ := db.Query("SELECT COUNT(*) FROM uploadfile")
	many := 0

	for m.Next() {
		m.Scan(&many)
	}

	return many
}
func GetFilesName(min, size int, filesChan chan []stuff.File) {
	q := `SELECT id,name FROM uploadfile WHERE  rowid >=?1 AND  rowid <=?2 ORDER BY id DESC ;` // get the filename and other stuff

	db := getConnection()
	defer db.Close()
	log.Println("starting")
	rows, err := db.Query(q, (size - (min * howMany)), (size-(min*howMany)+howMany)+1)

	if err != nil {
		log.Println(err)
		close(filesChan)
	}
	var filesList []stuff.File
	for rows.Next() {
		var file stuff.File
		if err := rows.Scan(&file.Id, &file.Name); err != nil {
			log.Println(err)
			close(filesChan)
		}

		filesList = append(filesList, file)

	}

	filesChan <- filesList
}
