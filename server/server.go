package server

import (
  "botota/database"
  "net/http"
  "fmt"
  "botota/APIHandler"
)

const (
  PORT = "3000"
)

func StartServer() {
  database.InitDB()

  http.HandleFunc("/", defaultHandler)
  http.ListenAndServe(fmt.Sprintf(":" + PORT), nil)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Hello")
}
