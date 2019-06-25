package main

import (
	"log"
	"net/http"
	"os"
)

var port string

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	if port = os.Getenv("PORT"); len(port) == 0 {
		log.Println("PORT not set. Defaulting to 3000")
		port = "3000"
	}
	log.Println("App is up and running at port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Println("ListenAndServe: ", err)
	}
}
