package main

import (
	"log"
	"net/http"

	"github.com/rerokai/web-diary/backend/handlers"
)

func main() {

	http.Handle("/frontend/", http.StripPrefix("/frontend/", http.FileServer(http.Dir("../frontend"))))

	http.HandleFunc("/today", handlers.TodayHandler)

	log.Println("Serv running")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
