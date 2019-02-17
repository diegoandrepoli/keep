package main

import (
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)


//list of keep
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
    router.HandleFunc("/keep/{id}", deleteKeepById).Methods("DELETE")

    fmt.Println("Running keep server!")
    log.Fatal(http.ListenAndServe(":8080", router))
}

/**
 * Main keep data
 */
func index(w http.ResponseWriter, r *http.Request){
   fmt.Fprintln(w, getIndexMessage())
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
    params := mux.Vars(r)
    for _, item := range keeps {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }

    json.NewEncoder(w).Encode(&Keep{})
}

/**
 * Post keep
 */
func postKeep(w http.ResponseWriter, r *http.Request) {
    var keep Keep
    _ = json.NewDecoder(r.Body).Decode(&keep)
    keeps = append(keeps, keep)
    json.NewEncoder(w).Encode(keep)
}


func deleteKeepById(w http.ResponseWriter, r *http.Request){
   params := mux.Vars(r)
   for index, item := range keeps {
      if item.ID == params["id"] {
         keeps = append(keeps[:index], keeps[index+1:]...)
         break
      }
    }
}
