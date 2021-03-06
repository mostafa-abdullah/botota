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
	byeMessage = "Thank you for using botota! It's been nice chatting with you!"
)

type Response struct {
	Messages []models.Message `json:"messages"`
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

func getReply(u models.User, msg string) []models.Message {
	reply := []models.Message{}
	q, err, done := info.Process(u, msg)

	if err != nil {
		highlight := err.Error()
		value := q.Text

		if q.Id == 5 {
			value += "\n" + formattedHotels(u.Hotels)
		}
		msg	:= models.Message{Highlight : highlight, Value: value}
		reply = append(reply, msg)
	}

	switch q.Id {
	case 0:
		if done {
			msg := models.Message{Highlight : byeMessage}
			reply = append(reply, msg)
			return reply
		}
	case 5:
		// Retrieve the list of hotels from database if not already retrieved
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
			highlight := q.Text
			value			:= formattedHotels(hotels)
			msg 			:= models.Message{Highlight: highlight, Value: value}
			reply = append(reply, msg)
			return reply
		case 6:
			// Gathered all info; return the schedule
			updatedUser, _ := database.Mongo.GetUser(u.Uuid)
			schedule := GooglePlacesAPIHandler.CreateSchedule(u.Destination, updatedUser.ChosenHotel, u.StartDate, u.EndDate)
			reply = append(reply, schedule...)
		}
		reply = append(reply, models.Message{Value: q.Text})
		return reply
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
