package models

type User struct {
  Uuid string
  Destination string
  StartDate string
  EndDate string
  Budget int
  Hotels []Place
  ChosenHotel Place
  CurrentQuestionId int
}
