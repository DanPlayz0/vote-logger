package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Would-You-Bot/vote-logger/botlists"
	"github.com/Would-You-Bot/vote-logger/config"
	"github.com/gorilla/mux"
)


func main() {
	config.Parse() 
	
	router := mux.NewRouter()
	router.HandleFunc("/topgg", botlists.HandleTopgg).Methods("POST")
	router.HandleFunc("/dscbot", botlists.HandleTopgg).Methods("POST")
	router.HandleFunc("/dlist", botlists.HandleDlistgg).Methods("POST")

	fmt.Printf("Server started on port 3340\n")
	errListen := http.ListenAndServe(":3340", router)
	log.Fatal(errListen)
}
