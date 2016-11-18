package chat

import (
	"botota/GooglePlacesAPIHandler"
	"botota/database"
	"botota/models"
	"botota/user/info"
	"encoding/json"
	"net/http"
	"strconv"
)

const (
	welcomeMessage = "Welcome to Botota! Your customized trip planner!"
	byeMessage     = "Thank you for using botota! It's been nice chatting with you!"
)

type Response struct {
	Message string `json:"message"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed.", http.StatusMethodNotAllowed)
		return
	}

	uuid := r.Header.Get("Authorization")
	user, auth := getAuthorizedUser(uuid)

	if !auth {
		http.Error(w, "Invalid UUID.", http.StatusUnauthorized)
		return
	}

	var body map[string]interface{}

	json.NewDecoder(r.Body).Decode(&body)

	// get the message in the body
	_, msgFound := body["message"]

	if !msgFound {
		http.Error(w, "Missing message key in request body.", http.StatusBadRequest)
		return
	}

	msg := body["message"].(string)

	reply := getReply(user, msg)
	res := Response{reply}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func getReply(u models.User, msg string) string {
	q, err, done := info.Process(u, msg)
	if err != nil {
		ret := err.Error() + "\n" + q.Text

		if q.Id == 5 {
			ret += "\n" + formattedHotels(u.Hotels)
		}

		return ret
	}

	reply := ""
	switch q.Id {
	case 0:
		if done {
			reply = byeMessage
		}
	case 1:
		reply = welcomeMessage + " "
	case 6:
		// Gathered all info; return the schedule + restart question
		updatedUser, _ := database.Mongo.GetUser(u.Uuid)
		schedule := GooglePlacesAPIHandler.CreateSchedule(u.Destination, updatedUser.ChosenHotel, u.StartDate, u.EndDate)
		reply = schedule + "\n"
	case 5:
		// Retrieve the list of hotels if not already retrieved
		var hotels []models.Place
		if len(u.Hotels) > 0 {
			hotels = u.Hotels
		} else {
			hotels = GooglePlacesAPIHandler.GetHotels(u.Destination)
			updatedUser, _ := database.Mongo.GetUser(u.Uuid)
			updatedUser.Hotels = hotels

			database.Mongo.UpdateUser(updatedUser)
		}

		// return formatted list of hotels
		return q.Text + `
    ` + formattedHotels(hotels)
	}
	return reply + q.Text

}

//formattedHotels returns a list of hotels in this form:
//1) Hotel1
//2) Hotel2
//....
func formattedHotels(hotels []models.Place) string {
	res := ""

	for i, h := range hotels {
		res += strconv.Itoa(i+1) + ") " + h.Name + `.`
	}

	return res
}

//getAuthorizedUser returns the user if he is authorized, otherwise returns a false flag
func getAuthorizedUser(uuid string) (models.User, bool) {
	user, exists := database.Mongo.GetUser(uuid)

	return user, exists
}
