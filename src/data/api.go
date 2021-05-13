package data

import "github.com/ranon-rat/simpleCloudInGO/src/stuff"

func GetFilesName(min int, apiChan chan []stuff.Api) {
	q := `SELECT (id,name) FROM publ 
	WHERE  rowid >=?1 AND  rowid <=?2
	ORDER BY id DESC ;`
	db := getConnection()
	size := GetSize()
	rows, _ := db.Query(q, (size - (min * howMany)), (size-(min*howMany)+howMany)+1) // envia esto y la salida deb de ser la siguiente
	var apiList []stuff.Api
	for rows.Next() {
		var api stuff.Api
		rows.Scan(&api)
		apiList = append(apiList, api)

	}
	apiChan <- apiList
}
