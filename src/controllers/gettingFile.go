package controllers

import (
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ranon-rat/simpleCloudInGO/src/data"
	"golang.org/x/sync/errgroup"
)

func GetFile(w http.ResponseWriter, r *http.Request) {
	name, fileChan, g := mux.Vars(r)["name"], make(chan io.ReadSeeker), new(errgroup.Group)
	g.Go(func() error { return data.Exist(name) })
	if g.Wait() == nil { // if the file exist return an error for that we are doing this
		http.Error(w, fmt.Sprint("file:", `"name"`, " doesnt find"), 404)
	}

	g.Go(func() error { return data.GetFile(name, fileChan) })
	if err := g.Wait(); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", 500)
		// if we get an error its going to be from the server so we are checking that
	}
	// Im using this for fomart the attachments and others stuff for make me more easy the work
	cd := mime.FormatMediaType("attachment", map[string]string{"filename": name})
	w.Header().Set("Content-Disposition", cd)
	// send the file for download
	http.ServeContent(w, r, name, time.Now(), <-fileChan)
}
