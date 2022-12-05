package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	log.Println("Listening on 80")
	err := http.ListenAndServe("0.0.0.0:80", nil)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
