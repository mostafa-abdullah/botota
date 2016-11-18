package server

import (
  "botota/api/welcome"
  "botota/api/chat"
  "botota/database"
  "net/http"
  "fmt"
  cors "github.com/heppu/simple-cors"
)

const (
  PORT = "3000"
)

func StartServer() {
  database.InitDB()

  mux := http.NewServeMux()
	mux.HandleFunc("/welcome", welcome.Handler)
	mux.HandleFunc("/chat", chat.Handler)
	mux.HandleFunc("/", defaultHandler)

  http.ListenAndServe(fmt.Sprintf(":" + PORT), cors.CORS(mux))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Hello")
}
