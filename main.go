package main

import (
    "fmt"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

/**
 * Main server
 */
func main() {
    var router = mux.NewRouter()
    router.HandleFunc("/keep", postKeep).Methods("POST")
    router.HandleFunc("/keep", handleKeep).Methods("GET")
    router.HandleFunc("/keep/{id}", handleKeepById).Methods("GET")

    fmt.Println("Running keep server!")
    log.Fatal(http.ListenAndServe(":8080", router))
}

/**
 * Handle all keeps
 */
func handleKeep(w http.ResponseWriter, r *http.Request) {
    //implement here
}

/**
 * Handle keep by identification
 */
func handleKeepById(w http.ResponseWriter, r *http.Request) {
    //implement here
}

/**
 * Post keep
 */
func postKeep(w http.ResponseWriter, r *http.Request) {
    //implement here
}

