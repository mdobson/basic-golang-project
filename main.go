package main

import (
  "log"
  "flag"
  "net/http"
  "github.com/gorilla/mux"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
  r := mux.NewRouter().StrictSlash(true)

  //handle root
  r.HandleFunc("/", HomeHandler).
  Methods("GET")


  r.HandleFunc("/players/{id}", ShowPlayerHandler).
  Methods("GET")

  r.HandleFunc("/market", ShowCompaniesHandler).
  Methods("GET")

  //handle create
  //r.HandleFunc("/create", CreateHandler).
  //Methods("POST")



  log.Fatal(http.ListenAndServe(*addr, r))
}
