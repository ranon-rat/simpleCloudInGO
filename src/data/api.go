package data

import (
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
func GetFilesName(min int, filesChan chan []stuff.File) {
	q := `SELECT id,name FROM uploadfile
	WHERE  rowid >=?1 AND  rowid <=?2
	ORDER BY id DESC ;` // get the filename and other stuff

	db := getConnection()
	defer db.Close()

	size := GetSize()

	rows, _ := db.Query(q, (size - (min * howMany)), (size-(min*howMany)+howMany)+1)

	var filesList []stuff.File

	for rows.Next() {
		name := ""
		id := 0
		rows.Scan(&id, &name)
		filesList = append(filesList, stuff.File{
			Name: name, Id: id,
		})

	}

	filesChan <- filesList
}
