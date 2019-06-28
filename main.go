package main

import (
	"log"
	"net/http"
	"os"
)

var port string

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{ "status": "pass" }`))
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	http.HandleFunc("/health", handleHealthCheck)

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
