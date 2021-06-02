package data

import (
	"github.com/ranon-rat/simpleCloudInGO/src/interfaces"
)

func GetSize() int {
	db := getConnection()
	defer db.Close()

	many := 0
	db.QueryRow("SELECT COUNT(*) FROM uploadfile").Scan(&many)

	return many
}

func GetFilesName(min, size int, filesChan chan []interfaces.File) {
	q := `SELECT id,name FROM uploadfile WHERE  rowid >=?1 AND  rowid <=?2 ORDER BY id DESC ;` // get the filename and other stuff

	db := getConnection()
	defer db.Close()
	rows, _ := db.Query(q, (size - (min * howMany)), (size-(min*howMany)+howMany)+1)

	var filesList []interfaces.File
	for rows.Next() {
		var file interfaces.File
		rows.Scan(&file.Id, &file.Name)

		filesList = append(filesList, file)

	}

	filesChan <- filesList
}
