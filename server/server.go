package server

import (
  "botota/api/welcome"
  "botota/api/chat"
  "botota/database"
  "net/http"
  "fmt"
)

const (
  PORT = "3000"
)

func StartServer() {
  database.InitDB()
  http.HandleFunc("/welcome", welcome.Handler)
  http.HandleFunc("/chat", chat.Handler)
  http.HandleFunc("/", defaultHandler)
  http.ListenAndServe(fmt.Sprintf(":" + PORT), nil)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Hello")
}
