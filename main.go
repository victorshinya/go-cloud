package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/victorshinya/go-cloud/internal/health"
)

var port string

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	http.HandleFunc("/health", health.HealthCheckHandler)

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
