package APIHandler

import (
  "net/http"
  "github.com/satori/go.uuid"
  "encoding/json"
  "botota/database"
  "botota/models"
)

type Response struct {
  Message string `json:"message"`
  Uuid string     `json:"uuid"`
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request){
  //create user uuid
  u := createUserUUID();

  //create a new User model
  user := models.User{Uuid: u}

  //insert the new user to the database
  database.Mongo.CreateUser(user);

  //get first question
  q := database.Mongo.GetFirstQuestion().Text;

  //prepare response
  res := Response{q, u}
  js, err := json.Marshal(res)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  //write json response
  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}

func createUserUUID() string {
  u := uuid.NewV4()
  uString := u.String();

  return uString;
}
