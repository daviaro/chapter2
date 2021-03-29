package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func main() {
	router := httprouter.New()
	router.ServeFiles("/static/*filepath", http.Dir("/Users/leydivivianarodriguezsanchez/workspace/src/github.com/daviaro/chapter2/fileServer/store"))
	log.Fatal(http.ListenAndServe(":8000", router))
}
