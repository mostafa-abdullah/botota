package welcome

import (
	"botota/database"
	"botota/models"
	"encoding/json"
	"github.com/satori/go.uuid"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Uuid    string `json:"uuid"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed.", http.StatusMethodNotAllowed)
		return
	}
	//create user uuid
	u := createUserUUID()

	//create a new User model
	user := models.User{Uuid: u}

	//insert the new user to the database
	database.Mongo.CreateUser(user)

	//get first question
	q := database.Mongo.GetFirstQuestion().Text

	//prepare response
	res := Response{q, u}

	//write json response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func createUserUUID() string {
	u := uuid.NewV4()
	uString := u.String()

	return uString
}
