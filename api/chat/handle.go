package chat

import (
  "net/http"
  "botota/models"
  "botota/database"
  "botota/user/info"
  "botota/GooglePlacesAPIHandler"
  "strconv"
  "encoding/json"
)

type Response struct {
  Message string `json:"message"`
}

func Handler(w http.ResponseWriter, r *http.Request){
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

  r.ParseForm()
  msgs := r.Form["message"]

  if len(msgs) == 0 {
    http.Error(w, "Missing message key in request body.", http.StatusBadRequest)
    return
  }

  // get the message in the body
  msg := msgs[0]

  reply := getReply(user, msg)
  res := Response{reply}

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(res)
}

func getReply(u models.User, msg string) string {
  q, err, done := info.Process(u, msg)
  if err != nil {
    return err.Error() + "\n" + q.Text;
  }

  if done {
    // Gathered all info; return the schedule
    return GooglePlacesAPIHandler.CreateSchedule(u.Destination, u.ChosenHotel, u.StartDate, u.EndDate)
  }

  if q.Id == 5 {
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

  return q.Text
}

//formattedHotels returns a list of hotels in this form:
//1) Hotel1
//2) Hotel2
//....
func formattedHotels(hotels []models.Place) string {
  res := ""

  for i, h := range hotels {
    res += strconv.Itoa(i+1) + ") " + h.Name + `.
    `
  }

  return res
}

//getAuthorizedUser returns the user if he is authorized, otherwise returns a false flag
func getAuthorizedUser(uuid string) (models.User, bool) {
  user, exists := database.Mongo.GetUser(uuid)

  return user, exists
}
