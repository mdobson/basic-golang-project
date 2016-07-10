package main

import (
  "encoding/json"
  "net/http"
  "github.com/gorilla/mux"
  "strconv"
)

type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

  root := CreateRootModel(r)
  js, err := json.Marshal(root)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/vnd.siren+json")
  w.Write(js)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {

}

func ShowPlayerHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id := vars["id"]
  var err error
  if playerId, err = strconv.Atoi(id); err != nil {
    panic(err)
  }

  player := StateFindPlayer(playerId)

  if player.id > 0 {
    playerModel := CreatePlayerModel(player, r)
    js, err := json.Marshal(playerModel)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    w.Header().Set("Content-Type", "application/vnd.siren+json")
    w.WriteHeader(http.StatusOK)
    w.Write(js)

    return
  }

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
  if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
    panic(err)
  }
}

func ShowCompaniesHandler(w http.ResponseWriter, r *http.Request) {
  market := StateFindAllCompanies()
  list := CreateCompanyList(market, r)
  js, err := json.Marshal(list)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/vnd.siren+json")
  w.Write(js)
}

func ShowCompanyHandler(w http.ResponseWriter, r *http.Request) {

}
