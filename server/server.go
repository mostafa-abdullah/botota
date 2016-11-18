package server

import (
  "botota/api/welcome"
  "botota/api/chat"
  "botota/database"
  "net/http"
  "fmt"
  "os"
  cors "github.com/heppu/simple-cors"
)

func StartServer() {
  database.InitDB()

  mux := http.NewServeMux()
	mux.HandleFunc("/welcome", welcome.Handler)
	mux.HandleFunc("/chat", chat.Handler)
	mux.HandleFunc("/", defaultHandler)

  port := os.Getenv("PORT")

  if port == "" {
    port = "3000"
  }

  http.ListenAndServe(fmt.Sprintf(":" + port), cors.CORS(mux))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Hello")
}
