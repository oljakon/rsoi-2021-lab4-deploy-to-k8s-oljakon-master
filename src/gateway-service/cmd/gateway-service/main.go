package main

import (
	"log"
	"net/http"
	"os"

	"rsoi2/src/gateway-service/internal/handlers"
)

func main() {
	port := os.Getenv("PORT")

	r := handlers.Router()

	log.Println("server is listening on port: ", port)
	log.Printf("app started")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
