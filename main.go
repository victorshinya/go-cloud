package main

import (
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var port string

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	http.HandleFunc("/health", HealthCheckHandler)

	if port = os.Getenv("PORT"); len(port) == 0 {
		log.Println("PORT not set. Defaulting to 3000")
		port = "3000"
	}
	log.Println("App is up and running at port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}

// HealthCheckHandler returns a status message whether the server is up or down
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{ "status": "UP" }`)
}
