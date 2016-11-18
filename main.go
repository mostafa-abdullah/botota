package main

import (
  "botota/server"
  "github.com/joho/godotenv"
)

func main(){
  godotenv.Load()
  server.StartServer()
}
