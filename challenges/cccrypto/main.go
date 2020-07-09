package main

import (
	"log"
	"net/http"
)

func submit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		r.ParseForm()

		if r.FormValue("username") == "796952" && r.FormValue("creditcardnumber") == "4971 9660 3200 7972" {
			log.Println("Correct flag submitted by: ", r.RemoteAddr)
			w.Write([]byte("Flag: CTF{15f5d85dd1fe2c86a6d04f96328b16b3}"))
		} else {
			log.Println("Incorrect flag submitted by:", r.RemoteAddr)
			w.Write([]byte("Nope"))
		}
	}
}

func main() {
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	http.HandleFunc("/submit", submit)
	log.Println("Listening on 3100")
	err := http.ListenAndServe("0.0.0.0:3100", nil)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
