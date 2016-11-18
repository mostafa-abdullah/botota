package info

import (
  "botota/models"
  "botota/database"
  "botota/utils"
  "errors"
  "regexp"
  "strconv"
  "time"
)

//GetInfo validates the message sent by the user and moves to next question
//receives the chatting user and the sent message
//returns the next question, an error message and a boolean to indicate if all questions were asked
func Process(u models.User, msg string) (models.Question, error, bool) {
  var nq models.Question

  if(u.CurrentQuestionId == 0) {
    // done with all the questions
    return nq, nil, true
  }


  // Get the question previously asked to the user
  cq := database.Mongo.GetCurrentQuestion(u)

  err := validate(cq, u, msg)

  if err != nil {
    return cq, err, false
  }

  // Validation succeeded. Store in the database
  u.CurrentQuestionId = cq.NextQuestionId
  storeInfo(u, msg, cq.Id)

  if(u.CurrentQuestionId == 0) {
    // done with all the questions
    return nq, err, true
  }

  nq = database.Mongo.GetNextQuestion(cq)

  return nq, err, false
}

//validate makes sure that the user sent a valid message according to the asked question.
func validate(cq models.Question, u models.User, msg string) error {
  reg, _ := regexp.Compile(cq.Regex)

  if !reg.MatchString(msg) {
      return errors.New("Invalid Input: the given message doesn't match the required format!")
  }

  //validate dates
  if cq.Id == 2 || cq.Id == 3 {
    form := "02/01/2006" //equivalent to dd/mm/yyyy
  	t, _ := time.Parse(form, msg)

    //validate future date
    if !t.After(time.Now()) {
      return errors.New("Invalid Input: The date should be in the future!")
    }

    //validate end date > start date
    if cq.Id == 3 {
      start, _  := time.Parse(form, u.StartDate)
      if !t.After(start){
        return errors.New("Invalid Input: Trip end date should be after its start date .")
      }
    }
  }

  if cq.Id == 5 {
    i, err := strconv.Atoi(msg)
    utils.Check(err)
    if  i > len(u.Hotels) || i <= 0 {
      return errors.New("Invalid Input: The hotel number is invalid.")
    }
  }

  return nil
}

//storeInfo stores the valid information sent by the user into the database.
func storeInfo(u models.User, msg string, qid int) {
  var err error

  switch qid {
  case 1:
    u.Destination = msg
  case 2:
    u.StartDate = msg
  case 3:
    u.EndDate = msg
  case 4:
    u.Budget, err = strconv.Atoi(msg)
    utils.Check(err)
  case 5:
    var hotelIdx int
    hotelIdx, err = strconv.Atoi(msg)
    utils.Check(err)
    u.ChosenHotel = u.Hotels[hotelIdx - 1]
  }

  database.Mongo.UpdateUser(u)
}
