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
	router.HandleFunc("/dscbot", botlists.HandleDscbot).Methods("POST")

	fmt.Printf("Server started on port %s\n", config.Conf.Port)
	errListen := http.ListenAndServe(":" + config.Conf.Port, router)
	log.Fatal(errListen)
}
