package server

import (
  "net/http"
  "fmt"
  "botota/APIHandler"
)

const (
  PORT = "3000"
)

func StartServer() {
  http.HandleFunc("/welcome", APIHandler.WelcomeHandler)
  http.HandleFunc("/", defaultHandler)
  http.ListenAndServe(fmt.Sprintf(":" + PORT), nil)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Hello")
}
