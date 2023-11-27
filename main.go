package main

import (
	"log"
	"mycloud/model/repository"
	"mycloud/router"
	"net/http"
	"time"
)

func test() {

}

func main() {

	defer repository.Db.Close()
	r := router.NewRouter()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r.Index(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

}
