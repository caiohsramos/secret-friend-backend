package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	InitDB()
	router := NewRouter()
	log.Println("Server up and running")
	log.Fatal(http.ListenAndServe(":3001", router))
}
