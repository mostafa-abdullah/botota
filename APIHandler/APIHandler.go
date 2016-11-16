package APIHandler
import (
    "net/http"
    "github.com/satori/go.uuid"
    "botota/models"
    "botota/database"
)
func WelcomeHandler(w http.ResponseWriter, r *http.Request){
  //create user uuid
  u := createUserUUID();

  //create a new User model
  user := models.User{Uuid: u}

  //insert the new user to the database
  database.CreateUser(user);

  //get first question
  q := database.GetFirstQuestion().Text;

  //create response
  res := {"message": q, "uuid": u}
  w.Header().Set("Content-Type", "application/json")
  w.Write(res);

}
func createUserUUID() string {
  u := uuid.NewV4()
  uString := u.String();
  // fmt.Printf("%s\n",uString);
  return uString;
}
