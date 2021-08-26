package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
    //Initialize router using mux
    router := mux.NewRouter()

    // Create the routes, we will be creating each function in future
    router.HandleFunc("/api/spots", getSpots).Methods("GET")
    router.HandleFunc("/api/spots/{id}", getSpot).Methods("GET")

    // Initialize a server, log.Fatal will throw an error if anything goes wrong
    log.Fatal(http.ListenAndServe(":8000", router))
}

func getSpots(res http.ResponseWriter, req *http.Request) {
	content, err := ioutil.ReadFile("data.json")    
    if err != nil {
        fmt.Println("Err")
    }
    fmt.Fprintf(res, string(content)) 
}



