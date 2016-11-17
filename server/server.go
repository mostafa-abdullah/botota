package server

import (
  "botota/database"
  "botota/APIHandler"
  "net/http"
  "fmt"
)

const (
  PORT = "3000"
)

func StartServer() {
  database.InitDB()
  http.HandleFunc("/welcome", APIHandler.WelcomeHandler)
  http.HandleFunc("/", defaultHandler)
  http.ListenAndServe(fmt.Sprintf(":" + PORT), nil)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Hello")
}
