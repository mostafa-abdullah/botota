package APIHandler
import (
    "fmt"
    "net/http"
    "github.com/satori/go.uuid"
)
func WelcomeHandler(w http.ResponseWriter, r *http.Request){
  u := createUserUUID();
  
}
func createUserUUID() string {
  u := uuid.NewV4()
  uString := u.String();
  // fmt.Printf("%s\n",uString);
  return uString;
}
