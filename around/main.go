package main

import (
    "fmt"
    "log"
    "net/http"   
	//import the libiary for different handlers for different http methods like what we have in Jupiter
	"github.com/gorilla/mux"
)

func main() {
    fmt.Println("started-service")

    r := mux.NewRouter()
    r.Handle("/upload", http.HandlerFunc(uploadHandler)).Methods("POST", "OPTIONS")
    log.Fatal(http.ListenAndServe(":8080", r))
}
