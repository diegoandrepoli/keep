package main

import (
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)


// The keep type
type Keep struct {
    ID       string   `json:"id,omitempty"`
    Title    string   `json:"title,omitempty"`
    Message  string   `json:"message,omitempty"`
}

var keeps []Keep



/**
 * Main server
 */
func main() {
    var router = mux.NewRouter()

    //keep demo api
    keeps = append(keeps, Keep{ID: "1", Title: "My keep", Message: "Keep display message"})
    keeps = append(keeps, Keep{ID: "2", Title: "My second keep", Message: "Keep display on second keep message"})

    router.HandleFunc("/", index).Methods("GET")
    router.HandleFunc("/keep", postKeep).Methods("POST")
    router.HandleFunc("/keep", handleKeep).Methods("GET")
    router.HandleFunc("/keep/{id}", handleKeepById).Methods("GET")

    fmt.Println("Running keep server!")
    log.Fatal(http.ListenAndServe(":8080", router))
}

/**
 * Main keep data
 */
func index(w http.ResponseWriter, r *http.Request){
   fmt.Fprintln(w, "Welcome to Keep!")
}

/**
 * Encode and response all keeps
 */
func handleKeep(w http.ResponseWriter, r *http.Request) {
   json.NewEncoder(w).Encode(keeps)
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

