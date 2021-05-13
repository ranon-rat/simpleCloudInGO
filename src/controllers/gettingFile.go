package controllers

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/ranon-rat/simpleCloudInGO/src/data"
	"golang.org/x/sync/errgroup"
)

func GetFile(w http.ResponseWriter, r *http.Request) {
	g, fileChan := new(errgroup.Group), make(chan io.ReadSeeker)
	idStr, name := mux.Vars(r)["id"], mux.Vars(r)["name"]
	id, _ := strconv.Atoi(idStr)

	g.Go(func() error { return data.Exist(id) })
	if g.Wait() == nil { // if the file exist return an error for that we are doing this
		http.Error(w, fmt.Sprint("file:", idStr, " doesnt find"), 404)
		return
	}

	go data.GetFile(id, fileChan)

	// Im using this for fomart the attachments and others stuff for make me more easy the work
	cd := mime.FormatMediaType("attachment", map[string]string{"filename": name[:len(name)-1]})
	w.Header().Set("Content-Disposition", cd)
	// send the file for download
	http.ServeContent(w, r, name, time.Now(), <-fileChan)
}
