package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"messaggioService/api"
	"messaggioService/db"
	"messaggioService/kafka"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db.InitDB()
	defer db.CloseDB()

	kafka.InitProducer()
	defer kafka.CloseProducer()

	go kafka.ConsumeMessages()

	router := mux.NewRouter()
	router.HandleFunc("/message", api.CreateMessage).Methods("POST")
	router.HandleFunc("/stats", api.GetStats).Methods("GET")

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
