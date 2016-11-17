package models

type Question struct {
  Id int
  Text string
  NextQuestionId int
  Regex string
}
